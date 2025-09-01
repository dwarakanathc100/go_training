# Day 2 — Git & GitHub Practicals (Step‑by‑Step)

**Goal:** By the end of this session you will create a GitHub account, set up Git on your computer, create a repository, make your first commit, and push it to GitHub.

> Recommended approach in class: **Everyone shares screen while doing the steps**. Trainer watches and unblocks quickly.


---

## ✅ Learning Outcomes
- Understand the difference between **Git (local VCS)** and **GitHub (remote hosting)**.
- Configure Git with your name/email.
- Authenticate to GitHub (SSH recommended; HTTPS + PAT optional).
- Create a **remote repository**, link it to a local repo, and **push**.
- Know the **daily workflow** (pull → branch → commit → push → PR → merge).


---

## 🧰 Prerequisites
- **Git installed** on your machine.  
  - Windows: install **Git for Windows** (comes with *Git Bash*).
  - macOS: `xcode-select --install` or install via Homebrew.
  - Linux (Debian/Ubuntu): `sudo apt update && sudo apt install git -y`
- A **GitHub account**: https://github.com (free).


---

## 1) Create Your GitHub Account (once)
1. Go to **https://github.com** → **Sign up**.
2. Choose a clear **username** (e.g., `firstname-lastname`).
3. Verify email and sign in.
4. Optional: Add a profile picture and your name.


---

## 2) Configure Git Locally (once per machine)
Open **Terminal / Git Bash / PowerShell** and run:

```bash
# Your real name & the same email you used on GitHub
git config --global user.name  "Your Name"
git config --global user.email "your-email@example.com"

# Show default config
git config --list --show-origin

# Make 'main' the default initial branch name
git config --global init.defaultBranch main
```


---

## 3) Authenticate to GitHub

### Option A — SSH (✅ Recommended)
SSH lets you push without typing a password every time.

```bash
# 1) Generate a key (use your GitHub email)
ssh-keygen -t ed25519 -C "your-email@example.com"

# Press Enter to accept default path (id_ed25519). Optionally set a passphrase.

# 2) Start the SSH agent and add your key
# macOS / Linux
eval "$(ssh-agent -s)"
ssh-add ~/.ssh/id_ed25519

# Windows (Git Bash)
eval $(ssh-agent -s)
ssh-add ~/.ssh/id_ed25519

# 3) Copy your public key to clipboard
# macOS
pbcopy < ~/.ssh/id_ed25519.pub
# Linux
xclip -sel clip < ~/.ssh/id_ed25519.pub   # install xclip if needed
# Windows (Git Bash)
cat ~/.ssh/id_ed25519.pub | clip
```

Now add the key to GitHub:  
**GitHub → Settings → SSH and GPG keys → New SSH key** → paste the key → Save.

Test the connection:

```bash
ssh -T git@github.com
# You should see a success message the first time.
```

> If prompted with authenticity prompt, type `yes` and press Enter.


### Option B — HTTPS + Personal Access Token (PAT)
1. GitHub → **Settings → Developer settings → Personal access tokens → Tokens (classic)**.  
2. **Generate new token** with scopes at least: `repo`. Copy it and store securely.
3. When you push for the first time, Git will ask for **username** and **password**.  
   - Enter GitHub username for username.  
   - **Paste the token** for password.  
4. To cache credentials on your machine:
   ```bash
   git config --global credential.helper store   # or 'manager' on Windows
   ```


---

## 4) Create a Remote Repository on GitHub
1. On GitHub, click **+** (top-right) → **New repository**.
2. Name it (e.g., `go-basics`). Keep **Public** for class.
3. **Important choice**:  
   - **Leave it empty** (no README) if you plan to **init locally and push** (simplest for beginners).  
   - If you **add a README** on GitHub now, you should **clone** instead of init locally to avoid conflicts.

We’ll use the **empty remote** path in class. Copy the **SSH URL**, e.g.:
```
git@github.com:YOUR-USERNAME/go-basics.git
```


---

## 5) Create a Project Locally & Make First Commit
```bash
# Create folder and init repo
mkdir go-basics && cd go-basics
git init  # or: git init -b main

# Add a README
echo "# Go Basics" > README.md

# OPTIONAL: Go-friendly .gitignore
cat > .gitignore << 'EOF'
# Go build caches and binaries
bin/
*.exe
*.exe~
*.dll
*.so
*.dylib
*.test
*.out

# Go coverage
coverage.*
*.coverprofile

# IDE / editor
.vscode/
.idea/
*.swp
.DS_Store
EOF

# Stage and commit
git add .
git commit -m "Initial commit: README and .gitignore"
```


---

## 6) Connect Local → Remote & Push
```bash
# Link the remote (use your SSH URL)
git remote add origin git@github.com:YOUR-USERNAME/go-basics.git

# Verify
git remote -v

# Push and set upstream
git push -u origin main
```

Go to your GitHub repo page and **refresh** — you should see your files. 🎉


---

## 7) Daily Workflow (Cheat Sheet)
```bash
# Get latest changes from main
git pull --ff-only origin main

# Create a feature branch
git checkout -b feature/hello-world

# Work, then stage & commit
git add .
git commit -m "Add hello world example"

# Push branch
git push -u origin feature/hello-world

# Open a Pull Request on GitHub, request review, and merge.
# After merge:
git checkout main
git pull --ff-only origin main
git branch -d feature/hello-world
git push origin --delete feature/hello-world   # optional
```


---

## 8) Clone an Existing Repository
If the repo already exists on GitHub (e.g., trainer’s starter code):

```bash
git clone git@github.com:TRAINER-USERNAME/repo-name.git
# or HTTPS
git clone https://github.com/TRAINER-USERNAME/repo-name.git
cd repo-name
```


---

## 9) Common Errors & Quick Fixes
**❌ Permission denied (publickey).**  
- SSH key not added to agent or GitHub.  
  ```bash
  ssh-add -l                      # list keys
  ls -la ~/.ssh                   # verify id_ed25519 and id_ed25519.pub
  ```

**❌ Push rejected (remote has work you don’t).**  
- You probably created a README on GitHub. Pull then push:
  ```bash
  git pull --rebase origin main
  git push origin main
  ```

**❌ Wrong author name/email in commits.**  
```bash
git config --global user.name  "Correct Name"
git config --global user.email "correct-email@example.com"
git commit --amend --reset-author
```

**❌ Asked for GitHub password over HTTPS.**  
- Use a **PAT** (token) instead of your password. See Option B above.


---

## 10) Optional: GitHub CLI (gh)
```bash
# macOS (Homebrew)
brew install gh

# Login
gh auth login

# Create a repo from terminal
gh repo create go-basics --public --source=. --remote=origin --push
```


---

## 11) Quick Git Command Map
- `git init` — make a folder a Git repo  
- `git status` — see what changed  
- `git add .` — stage changes  
- `git commit -m "msg"` — save a snapshot  
- `git log --oneline --graph --decorate` — nice history view  
- `git remote -v` — list remotes  
- `git push / git pull` — sync with GitHub  
- `git checkout -b <branch>` — new branch  
- `git switch <branch>` — move between branches  


---

## 12) In‑Class Mini Tasks
1) Create a repo named **`go-basics`** and push a README.  
2) Add a file **`hello.go`** and push it on a new branch `feature/hello`.  
3) Open a **Pull Request**, review with a partner, and merge.


---

### Appendix A — Install Git (quick)
- **Windows**: https://git-scm.com/download/win (choose default options; use *Git Bash*)
- **macOS**: `xcode-select --install`  or  `brew install git`
- **Ubuntu/Debian**: `sudo apt update && sudo apt install git -y`

### Appendix B — Sample `hello.go`
```go
package main

import "fmt"

func main() {
    fmt.Println("Hello from Go!")
}
```

---

**Tip:** Commit small, commit often. Write meaningful messages.  
**Mantra:** *Pull → Branch → Commit → Push → PR → Merge*.
