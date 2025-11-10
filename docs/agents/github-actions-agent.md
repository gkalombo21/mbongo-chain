# Cursor GitHub Actions Agent

Ce script PowerShell enregistre un agent léger dans Cursor qui surveille vos
commandes `git`. Lorsqu'un `git push` réussit sur la branche `main`, l'agent
ouvre automatiquement la page GitHub Actions du dépôt afin de visualiser le
workflow associé au nouveau commit.

## Installation

1. Assurez-vous d'utiliser un terminal PowerShell dans Cursor (Windows est
   supporté nativement).
2. Depuis la racine du dépôt, exécutez :

   ```powershell
   .\scripts\cursor-github-actions-agent.ps1
   ```

   Vous pouvez adapter les paramètres si nécessaire :

   ```powershell
   .\scripts\cursor-github-actions-agent.ps1 -BranchName main -RemoteName origin -RepositoryUrl https://github.com/gkalombo21/mbongo-chain.git
   ```

   - `BranchName` : branche à surveiller (`main` par défaut).
   - `RemoteName` : remote Git utilisé pour récupérer le commit distant (`origin` par défaut).
   - `RepositoryUrl` : URL GitHub du dépôt. Si omise, l'agent la déduit de `git remote get-url <remote>`.

## Fonctionnement

- Le script stocke l'état du dernier commit distant visible avec
  `git ls-remote <remote> <branch>`.
- Après chaque commande saisie dans le terminal, le prompt PowerShell appelle
  l'agent.
- Si la commande est un `git push` et que la branche active est celle
  surveillée, l'agent compare l'empreinte du commit distant.
- Si un nouveau hash est détecté, `Start-Process` ouvre l'URL
  `https://github.com/<owner>/<repo>/actions`.

Aucune boucle infinie n'est utilisée : les vérifications se produisent
uniquement lorsque vous lancez des commandes Git dans le terminal Cursor.

## Désactivation

Redémarrez le terminal ou réaffectez votre fonction `prompt` si vous souhaitez
désactiver l'agent au cours de la session actuelle. Les modifications sont
volatiles et propres à la session.

