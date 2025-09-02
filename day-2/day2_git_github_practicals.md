# Day 2 — Git & GitHub Practicals (Step‑by‑Step)

**Goal:** By the end of this session you will create a GitHub account, set up Git on your computer, create a repository, make your first commit, and push it to GitHub.

---


##  1. What is Git?
- Git is a **distributed version control system (VCS)**.
- It keeps track of **changes in code**, who made them, and when.
- Allows multiple developers to work on the same project **simultaneously** without overwriting each other’s work.

 Think of Git as a **time machine + collaboration tool** for your code.

---

##  2. Why Do We Need Git?
- **Code Sharing**: Collaborate easily with teammates.
- **Tracking Changes**: See what changed, when, and by whom.
- **Version Control**: Revert back if something breaks.
- **Collaboration**: Work on different features without conflicts.

---

##  3. Git Workflow Overview
1. **Working Directory** → where you edit files.
2. **Staging Area** → where you mark files to be committed.
3. **Repository (Local Repo)** → where Git stores committed versions.
4. **Remote Repository (GitHub, GitLab, etc.)** → shared version for the whole team.




##  Learning Outcomes
- Understand the difference between **Git (local VCS)** and **GitHub (remote hosting)**.
- Configure Git with your name/email.
- Authenticate to GitHub.
- Create a **remote repository**, link it to a local repo, and **push**.
- Know the **daily workflow** (pull → branch → commit → push → PR → merge).


---

##  Prerequisites
- **Git installed** on your machine.  
  - Windows: install **Git for Windows** 
  **Windows**: https://git-scm.com/download/win

- A **GitHub account**: https://github.com (free).


---

## 1) Create Your GitHub Account (once)
1. Go to **https://github.com** → **Sign up**.
2. Choose a clear **username** (e.g., `firstname-lastname`).
3. Verify email and sign in.


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
git config --global init.defaultBranch main/master
```


---

## 3) Authenticate to GitHub
### Option A  - Using UI
When you push for the first time, Git will take you to the github.com for authorization.




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



# Stage and commit
git add .
git commit -m "Initial commit: README"
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

## 7) Daily Workflow 
```bash
# Get latest changes from main
git pull 
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
git pull origin main
git branch -d feature/hello-world
git push origin --delete feature/hello-world  
```


---

## 8) Clone an Existing Repository
If the repo already exists on GitHub

```bash
git clone git@github.com:TRAINER-USERNAME/repo-name.git
# or HTTPS
git clone https://github.com/TRAINER-USERNAME/repo-name.git
cd repo-name
```


---


## 9) Quick Git Command Map
- `git init` — make a folder a Git repo  
- `git status` — see what changed  
- `git add .` — stage changes  
- `git commit -m "msg"` — save a snapshot  
- `git log --oneline --graph --decorate` — nice history view  
- `git push / git pull` — sync with GitHub  
- `git checkout -b <branch>` — new branch  
- `git switch <branch>` — move between branches  


---

## 10)  Exercises
1) Create a repo named **`go-basics`** and push a README.  
2) Add a file **`hello.go`** and push it on a new branch `feature/hello`.  
3) Open a **Pull Request**, review with a partner, and merge.


---



## Sample `hello.go`
```go
package main

import "fmt"

func main() {
    fmt.Println("Hello from Go!")
}
```

---
## Final message:
 Commit small, commit often. Write meaningful messages.  
 Pull → Branch → Commit → Push → PR → Merge.
 
 