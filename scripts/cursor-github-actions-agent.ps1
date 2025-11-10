<#
.SYNOPSIS
Registers a lightweight Cursor agent that watches Git commands and opens the
GitHub Actions page whenever a new commit is pushed to the main branch.

.DESCRIPTION
The agent hooks into the PowerShell prompt lifecycle (executed after every
command) to inspect recently executed commands. When it sees a Git command,
it keeps track of the remote HEAD of the configured branch. If the remote HEAD
changes—indicating a successful push—it launches the default browser on the
repository's GitHub Actions page.

No polling loops are used; checks happen only after Git commands finish.

.NOTES
Requires PowerShell with access to Git and a default browser launcher
(`Start-Process` on Windows).
#>
param(
    [string]$BranchName = "main",
    [string]$RemoteName = "origin",
    [string]$RepositoryUrl
)

Set-StrictMode -Version Latest

function Get-CursorAgentRepositoryUrl {
    param([string]$RemoteName, [string]$RepositoryUrl)

    if ($RepositoryUrl) {
        return $RepositoryUrl
    }

    $remote = git remote get-url $RemoteName 2>$null
    if (-not $remote) {
        throw "Impossible de déterminer l'URL du dépôt pour le remote '$RemoteName'."
    }

    return $remote
}

function ConvertTo-GitHubActionsUrl {
    param([string]$RepositoryUrl)

    $normalized = $RepositoryUrl.Trim()
    if ($normalized.EndsWith(".git")) {
        $normalized = $normalized.Substring(0, $normalized.Length - 4)
    }

    if ($normalized -notmatch "^https://github\.com/") {
        throw "Seules les URLs GitHub https://github.com/<owner>/<repo>(.git) sont supportées (reçu '$RepositoryUrl')."
    }

    return "$normalized/actions"
}

function Get-RemoteBranchHead {
    param([string]$RemoteName, [string]$BranchName)

    $result = git ls-remote --heads $RemoteName $BranchName 2>$null
    if (-not $result) {
        return $null
    }

    $hash = $result.Split("`t", "`n", [System.StringSplitOptions]::RemoveEmptyEntries)[0]
    return $hash
}

function Invoke-CursorAgentCheck {
    param(
        [string]$CommandLine,
        [hashtable]$State
    )

    if ($CommandLine -notmatch "^\s*git\b") {
        return
    }

    $State.LastGitCommand = $CommandLine

    if ($CommandLine -notmatch "^\s*git\s+push\b") {
        return
    }

    $currentBranch = git rev-parse --abbrev-ref HEAD 2>$null
    if (-not $currentBranch -or $currentBranch -ne $State.BranchName) {
        return
    }

    $remoteHead = Get-RemoteBranchHead -RemoteName $State.RemoteName -BranchName $State.BranchName
    if (-not $remoteHead) {
        return
    }

    if ($State.LastRemoteHead -and $remoteHead -eq $State.LastRemoteHead) {
        return
    }

    $State.LastRemoteHead = $remoteHead

    try {
        Start-Process $State.ActionsUrl | Out-Null
    }
    catch {
        Write-Warning "Impossible d'ouvrir la page GitHub Actions ($($State.ActionsUrl))."
    }
}

function Register-CursorPromptHook {
    param([hashtable]$State)

    if ($State.PromptHookRegistered) {
        return
    }

    $existingPrompt = (Get-Command prompt -ErrorAction SilentlyContinue)?.ScriptBlock
    $State.OriginalPrompt = $existingPrompt

    function global:prompt {
        $state = $script:CursorGitActionsAgentState

        if ($state) {
            $history = Get-History -Count 1 -ErrorAction SilentlyContinue
            if ($history -and ($state.LastHistoryId -ne $history.Id)) {
                $state.LastHistoryId = $history.Id
                Invoke-CursorAgentCheck -CommandLine $history.CommandLine -State $state
            }
        }

        if ($state -and $state.OriginalPrompt) {
            return & $state.OriginalPrompt
        }

        return "PS " + (Get-Location) + "> "
    }

    $State.PromptHookRegistered = $true
}

if (-not (git rev-parse --is-inside-work-tree 2>$null)) {
    throw "Ce script doit être exécuté depuis un dépôt Git."
}

$repoUrl = Get-CursorAgentRepositoryUrl -RemoteName $RemoteName -RepositoryUrl $RepositoryUrl
$actionsUrl = ConvertTo-GitHubActionsUrl -RepositoryUrl $repoUrl
$lastRemoteHead = Get-RemoteBranchHead -RemoteName $RemoteName -BranchName $BranchName
$history = Get-History -Count 1 -ErrorAction SilentlyContinue

$script:CursorGitActionsAgentState = @{
    BranchName          = $BranchName
    RemoteName          = $RemoteName
    RepositoryUrl       = $repoUrl
    ActionsUrl          = $actionsUrl
    LastRemoteHead      = $lastRemoteHead
    LastHistoryId       = $history?.Id
    PromptHookRegistered = $false
}

Register-CursorPromptHook -State $script:CursorGitActionsAgentState

Write-Verbose "Agent Cursor GitHub Actions enregistré pour $($actionsUrl)."

