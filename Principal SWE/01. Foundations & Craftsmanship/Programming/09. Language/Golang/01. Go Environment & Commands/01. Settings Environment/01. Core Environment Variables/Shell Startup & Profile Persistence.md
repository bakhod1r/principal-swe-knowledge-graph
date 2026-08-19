---
title: "Shell Startup & Profile Persistence"
tags:
  - review
  - golang
  - environment
  - principal-swe
parent: "[[Core Environment Variables]]"
---
# Shell Startup & Profile Persistence

**Shell startup** is the process by which a shell initializes its environment when you open a terminal or start a shell process.

**Profile persistence** means configuring shell settings—such as environment variables, aliases, functions, and `PATH` changes—so they are automatically available in future shell sessions.

## 1. Why does this exist?

A shell process starts with an environment inherited from its parent process. However, you often need to customize that environment:

```bash
export PATH="$HOME/bin:$PATH"
export GOPATH="$HOME/go"

alias ll='ls -lah'

function gs() {
    git status
}
```

If you execute these commands manually, they normally affect **only the current shell session**.

To make them persistent, you place them in a shell startup/profile file.

---

## 2. The key mental model

Think of shell initialization as:

```text
Terminal / Process
       │
       ▼
   Start shell
       │
       ▼
Determine shell mode
       │
       ▼
Read applicable startup files
       │
       ▼
Execute configuration
       │
       ▼
Interactive shell ready
```

The important part is:

> **Which startup file gets executed depends on the shell and how the shell was started.**

This is why blindly adding everything to `.bashrc` or `.profile` often causes confusing behavior.

---

# 3. Bash startup files

For Bash, the important files are:

|File|Typical purpose|
|---|---|
|`~/.bash_profile`|Login-shell configuration|
|`~/.bash_login`|Alternative login-shell file|
|`~/.profile`|Generic login/session configuration|
|`~/.bashrc`|Interactive Bash-shell configuration|
|`/etc/profile`|System-wide login configuration|
|`/etc/bash.bashrc`|System-wide interactive Bash configuration on many Linux distributions|

### Login shell

A login shell is associated with logging into the system.

For example:

```bash
ssh user@server
```

Bash may read:

```text
/etc/profile
~/.bash_profile
```

If `.bash_profile` does not exist, Bash checks other login startup files according to its rules.

---

### Interactive non-login shell

For example:

```bash
bash
```

An interactive Bash shell normally reads:

```text
~/.bashrc
```

This is where you typically put:

```bash
alias ll='ls -lah'
alias k='kubectl'

export PATH="$HOME/bin:$PATH"
```

---

# 4. `.bash_profile` vs `.bashrc`

This is one of the most important distinctions.

A common setup is:

```bash
~/.bash_profile
        │
        └── source ~/.bashrc
```

For example:

```bash
if [ -f ~/.bashrc ]; then
    . ~/.bashrc
fi
```

Then you can conceptually separate configuration:

```text
.profile / .bash_profile
        │
        ├── login/session environment
        │
        └── .bashrc
              │
              ├── aliases
              ├── functions
              ├── interactive settings
              └── shell customization
```

This prevents configuration from being duplicated.

---

# 5. `.profile`

`.profile` is more shell-agnostic than `.bashrc`.

It is commonly used for environment configuration such as:

```bash
export PATH="$HOME/bin:$PATH"
export EDITOR=vim
export GOPATH="$HOME/go"
```

The important distinction is:

```text
.profile
    → login/session environment

.bashrc
    → interactive Bash behavior
```

For example, an alias generally belongs in:

```bash
~/.bashrc
```

while a persistent environment variable may belong in:

```bash
~/.profile
```

depending on your OS/session architecture.

---

# 6. Zsh

If you use Zsh, the equivalent interactive configuration is usually:

```bash
~/.zshrc
```

For example:

```bash
export PATH="$HOME/bin:$PATH"

alias ll='ls -lah'
alias k='kubectl'
```

Zsh also has other startup files, including:

```text
~/.zprofile
~/.zshrc
~/.zlogin
~/.zlogout
```

A useful mental model is:

```text
Bash
├── .bash_profile
├── .bashrc
└── .profile

Zsh
├── .zprofile
├── .zshrc
├── .zlogin
└── .zlogout
```

---

# 7. Environment variable persistence

Suppose you run:

```bash
export APP_ENV=production
```

You can verify:

```bash
echo "$APP_ENV"
```

But after closing the terminal:

```bash
echo "$APP_ENV"
```

may produce nothing.

Why?

Because:

```text
Shell A
  │
  └── APP_ENV=production
```

When Shell A exits, its process environment disappears.

If you put:

```bash
export APP_ENV=production
```

inside the appropriate startup file:

```text
~/.bashrc
```

or:

```text
~/.profile
```

then every applicable new shell executes that configuration.

---

# 8. `export` is important

There is a subtle distinction:

```bash
APP_ENV=production
```

creates a shell variable.

Whereas:

```bash
export APP_ENV=production
```

creates an environment variable that child processes inherit.

For example:

```bash
export DATABASE_URL="postgres://localhost/app"
./my-program
```

The process tree becomes conceptually:

```text
Shell
 ├── DATABASE_URL
 │
 └── my-program
       └── inherits DATABASE_URL
```

Without `export`, the child process generally does not receive it.

---

# 9. PATH persistence

One of the most common profile modifications is:

```bash
export PATH="$HOME/bin:$PATH"
```

Suppose:

```bash
echo "$PATH"
```

returns:

```text
/usr/local/bin:/usr/bin:/bin
```

After:

```bash
export PATH="$HOME/bin:$PATH"
```

you get:

```text
/home/user/bin:/usr/local/bin:/usr/bin:/bin
```

Now executables inside:

```text
~/bin
```

can be found by commands such as:

```bash
mytool
```

without specifying:

```bash
~/bin/mytool
```

### Important principle

Do **not** casually overwrite `PATH`:

```bash
export PATH="$HOME/bin"
```

This can remove system directories and cause commands to stop working.

Prefer:

```bash
export PATH="$HOME/bin:$PATH"
```

or append deliberately:

```bash
export PATH="$PATH:$HOME/bin"
```

The order matters because the shell searches directories from left to right.

---

# 10. How to inspect what is actually happening

Instead of guessing which file Bash uses, inspect the shell.

```bash
echo "$SHELL"
```

For Bash:

```bash
bash --version
```

Determine whether the current shell is interactive:

```bash
case $- in
    *i*) echo "interactive" ;;
    *)   echo "non-interactive" ;;
esac
```

Determine whether it is a login shell:

```bash
shopt -q login_shell && echo "login" || echo "non-login"
```

You can also inspect the process:

```bash
ps -p $$ -o pid,ppid,args
```

Here `$$` is the current shell's PID.

---

# 11. Testing configuration safely

After modifying:

```bash
~/.bashrc
```

you can reload it:

```bash
source ~/.bashrc
```

or:

```bash
. ~/.bashrc
```

Then verify:

```bash
echo "$PATH"
```

and:

```bash
type ll
```

For an environment variable:

```bash
echo "$APP_ENV"
```

For executable resolution:

```bash
command -v go
command -v kubectl
```

---

# 12. Common mistake: "It works in my terminal but not in my application"

This is a classic environment-boundary problem.

You may have:

```bash
export PATH="$HOME/go/bin:$PATH"
```

in:

```text
~/.bashrc
```

and therefore:

```bash
go
```

works in your terminal.

But an IDE, system service, cron job, Docker process, or GUI application may **not execute `.bashrc`**.

Therefore:

```text
Terminal environment
        ≠
IDE environment
        ≠
systemd environment
        ≠
cron environment
        ≠
Docker environment
```

This is a crucial production engineering principle:

> **Shell configuration is not a universal environment-management mechanism.**

For production services, prefer explicit environment configuration through the service manager, container runtime, orchestration platform, or deployment system.

---

# 13. Shell startup vs system-wide configuration

There are two scopes:

### User-specific

```text
~/.bashrc
~/.profile
~/.zshrc
```

Only affects that user.

### System-wide

Examples:

```text
/etc/profile
/etc/environment
/etc/bash.bashrc
```

These affect broader system behavior.

For personal development configuration, prefer user-level configuration whenever possible.

For production infrastructure, avoid modifying global shell profiles as an application configuration mechanism.

---

# 14. Production perspective

A strong engineering distinction is:

```text
Shell configuration
        ↓
Developer convenience

Application configuration
        ↓
Explicit runtime configuration

Infrastructure configuration
        ↓
Deployment/system configuration
```

For example, this is reasonable in `.bashrc`:

```bash
alias k='kubectl'
alias ll='ls -lah'
```

But this is usually a poor production configuration strategy:

```bash
export DATABASE_PASSWORD="..."
export PAYMENT_API_KEY="..."
```

Secrets should be managed through appropriate secret-management mechanisms rather than shell profiles.

---

## Practical rule of thumb

```text
Alias / shell function
    → .bashrc / .zshrc

Interactive shell customization
    → .bashrc / .zshrc

Login/session environment
    → .profile / .bash_profile / .zprofile

Application configuration
    → application config / environment / deployment system

Production secrets
    → secret manager / orchestrator secrets

System service configuration
    → systemd / container / Kubernetes / deployment configuration
```

### Principal-level takeaway

The important concept is not memorizing `.bashrc` vs `.profile`.

It is understanding **process boundaries and initialization boundaries**:

```text
Parent process
     │
     │ inherits environment
     ▼
Shell
     │
     ├── startup files modify shell state
     │
     └── child processes inherit exported environment
```

Once you understand that model, shell startup behavior becomes much easier to debug.

---

## 🔗 References
- ⬆️ Parent: [[Core Environment Variables]]
- 📚 Module: `Go Environment & Commands`
