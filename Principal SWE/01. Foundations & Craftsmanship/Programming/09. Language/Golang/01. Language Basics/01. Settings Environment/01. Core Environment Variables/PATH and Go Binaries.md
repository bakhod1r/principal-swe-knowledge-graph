---
title: "PATH and Go Binaries"
tags:
  - review
  - golang
  - environment
  - principal-swe
parent: "[[Core Environment Variables]]"
---

# `PATH` and Go Binaries

This is a fundamental concept in Go development because tools such as `gopls`, `goimports`, `golangci-lint`, and custom Go CLIs are often installed as executables.

The core relationship is:

```text
go install
    ↓
Go binary
    ↓
GOBIN / GOPATH/bin
    ↓
PATH
    ↓
Shell can execute the binary
```

## 1. What is `PATH`?

`PATH` is an operating-system environment variable containing directories where the shell searches for executable programs.

Check it:

```bash
echo "$PATH"
```

Example:

```text
/usr/local/bin:/usr/bin:/bin:/home/user/go/bin
```

The `:` separates directories on Linux/macOS.

When you type:

```bash
gopls
```

the shell effectively searches:

```text
/usr/local/bin/gopls
/usr/bin/gopls
/bin/gopls
/home/user/go/bin/gopls
```

until it finds an executable.

---

# 2. What is a Go binary?

A Go program is compiled into an executable binary:

```bash
go build
```

For example:

```bash
go build -o myapp .
```

produces:

```text
myapp
```

You can execute it directly:

```bash
./myapp
```

Notice the `./`.

That's because the current directory is **usually not part of `PATH`**.

---

# 3. `go install` and binaries

The more interesting case is:

```bash
go install example.com/tool@latest
```

`go install` builds the executable and installs it into the configured binary directory.

Check:

```bash
go env GOBIN
```

If `GOBIN` is empty, the default location is generally:

```text
$(go env GOPATH)/bin
```

For example:

```text
/home/user/go/bin
```

So:

```bash
go install golang.org/x/tools/gopls@latest
```

might result in:

```text
~/go/bin/gopls
```

---

# 4. The critical distinction: `GOBIN` vs `PATH`

These two variables solve completely different problems.

### `GOBIN`

Answers:

> **Where should Go install my executable?**

```bash
go env GOBIN
```

### `PATH`

Answers:

> **Where should the shell look for executables?**

```bash
echo "$PATH"
```

Therefore:

```text
GOBIN
  ↓
~/go/bin/gopls

PATH
  ↓
must contain ~/go/bin
```

If `GOBIN` points to:

```text
/home/user/.local/bin
```

then:

```text
/home/user/.local/bin
```

must be in `PATH` if you want to execute the binary by name.

---

# 5. Recommended setup

For a standard Go installation, you usually don't need to configure `GOBIN`.

Check:

```bash
go env GOPATH
```

Suppose:

```text
/home/user/go
```

Then add:

```text
/home/user/go/bin
```

to your `PATH`.

A robust shell configuration is:

```bash
export PATH="$PATH:$(go env GOPATH)/bin"
```

For Bash:

```text
~/.bashrc
```

For Zsh:

```text
~/.zshrc
```

Then reload:

```bash
source ~/.bashrc
```

or:

```bash
source ~/.zshrc
```

---

# 6. Verify the complete chain

Install a tool:

```bash
go install golang.org/x/tools/gopls@latest
```

Find where Go installed it:

```bash
ls "$(go env GOPATH)/bin"
```

You should see something like:

```text
gopls
```

Now check whether the shell can find it:

```bash
command -v gopls
```

Expected:

```text
/home/user/go/bin/gopls
```

Finally:

```bash
gopls version
```

If all three work:

```text
go install
    ↓
~/go/bin/gopls
    ↓
~/go/bin in PATH
    ↓
command -v gopls
    ↓
gopls
```

your setup is correct.

---

# 7. Why `command -v` is better than `which`

You will often see:

```bash
which gopls
```

but for shell scripting and debugging, prefer:

```bash
command -v gopls
```

It is a shell builtin and gives you more reliable information about how the command is resolved.

For example:

```bash
command -v go
command -v gopls
command -v golangci-lint
```

---

# 8. Common failure

Suppose:

```bash
go install example.com/mytool@latest
```

succeeds.

But:

```bash
mytool
```

returns:

```text
command not found
```

Don't reinstall the tool immediately.

Debug systematically.

### Step 1 — Does the binary exist?

```bash
ls "$(go env GOPATH)/bin/mytool"
```

If it exists, installation succeeded.

### Step 2 — Is its directory in `PATH`?

```bash
echo "$PATH"
```

Or:

```bash
echo "$PATH" | tr ':' '\n'
```

Look for:

```text
/home/user/go/bin
```

### Step 3 — What does Go think the binary directory is?

```bash
go env GOBIN
go env GOPATH
```

If:

```text
GOBIN=""
GOPATH="/home/user/go"
```

the expected binary directory is:

```text
/home/user/go/bin
```

### Step 4 — Does the shell know about the updated `PATH`?

After changing your profile:

```bash
source ~/.bashrc
```

Then:

```bash
command -v mytool
```

---

# 9. Another subtle problem: stale shell command cache

Some shells cache command locations.

If you changed `PATH` and the shell still behaves unexpectedly, Bash can refresh its command hash:

```bash
hash -r
```

Then:

```bash
command -v mytool
```

This is particularly useful after installing or moving executables.

---

# 10. `PATH` ordering matters

Suppose you have:

```text
/usr/bin/tool
/home/user/go/bin/tool
```

and:

```bash
PATH="/usr/bin:/home/user/go/bin"
```

then:

```bash
tool
```

resolves to:

```text
/usr/bin/tool
```

because `/usr/bin` appears first.

Reverse it:

```bash
PATH="/home/user/go/bin:/usr/bin"
```

and the Go-installed version wins.

You can see what is being selected:

```bash
command -v tool
```

And sometimes:

```bash
type -a tool
```

is useful to find multiple matches.

---

# 11. Don't blindly modify `PATH`

A common bad configuration is:

```bash
export PATH="/home/user/go/bin"
```

This replaces the existing `PATH`.

You may suddenly lose access to:

```bash
ls
git
ssh
docker
go
```

depending on where they are installed.

Prefer:

```bash
export PATH="$PATH:$(go env GOPATH)/bin"
```

or, if you intentionally want the Go binaries to have precedence:

```bash
export PATH="$(go env GOPATH)/bin:$PATH"
```

These have different semantics because of ordering.

---

# 12. `GOBIN` can be useful

You can intentionally centralize user-installed executables:

```bash
go env -w GOBIN="$HOME/.local/bin"
```

Then:

```bash
go install golang.org/x/tools/gopls@latest
```

produces:

```text
~/.local/bin/gopls
```

And configure:

```bash
export PATH="$HOME/.local/bin:$PATH"
```

This gives you:

```text
Go tools
   ↓
~/.local/bin
   ↓
PATH
```

This can be a clean setup if you already use `~/.local/bin` for user-installed executables.

But don't introduce `GOBIN` just for the sake of customization. The default:

```text
$GOPATH/bin
```

is perfectly reasonable.

---

# 13. Project binaries are different

Consider:

```bash
go build -o bin/api ./cmd/api
```

You get:

```text
project/
└── bin/
    └── api
```

You can run it with:

```bash
./bin/api
```

You **do not need** to add `project/bin` to your global `PATH`.

This distinction is useful:

```text
go install
    → user/global developer tool
    → GOBIN / GOPATH/bin

go build
    → project artifact
    → chosen output directory

./bin/api
    → explicit project executable
```

For production applications, you generally build an explicit artifact and deploy it rather than relying on the developer's `PATH`.

---

# 14. Production perspective

`PATH` is primarily a **process environment concern**.

For example, your shell might have:

```text
PATH=/home/user/go/bin:/usr/local/bin:/usr/bin
```

But a systemd service may have a completely different environment.

Similarly:

```text
Developer shell
    ≠
CI runner
    ≠
Docker container
    ≠
Kubernetes Pod
    ≠
systemd service
```

Therefore, if your application only works because your personal `.bashrc` modifies `PATH`, the deployment is fragile.

A production build should explicitly establish its required environment.

---

# 15. A clean Go developer setup

For most Linux/macOS Go developers:

```bash
go version
go env GOPATH
```

Then:

```bash
export PATH="$PATH:$(go env GOPATH)/bin"
```

Persist that in the appropriate shell startup file.

Verify:

```bash
go install golang.org/x/tools/gopls@latest

command -v gopls
gopls version
```

The resulting architecture is simple:

```text
                    ┌──────────────┐
                    │ go install   │
                    └──────┬───────┘
                           │
                           ▼
                    ┌──────────────┐
                    │ GOBIN        │
                    │ or           │
                    │ GOPATH/bin   │
                    └──────┬───────┘
                           │
                           ▼
                    ┌──────────────┐
                    │ executable   │
                    └──────┬───────┘
                           │
                           ▼
                    ┌──────────────┐
                    │ PATH         │
                    └──────┬───────┘
                           │
                           ▼
                    ┌──────────────┐
                    │ shell        │
                    │ finds tool   │
                    └──────────────┘
```

### The key mental model

**`go install` determines where the binary goes; `PATH` determines whether you can invoke it by name.**

If a Go CLI is installed but `command not found` appears, first investigate **binary location → `GOBIN`/`GOPATH` → `PATH` → shell cache**. Don't reinstall blindly.

---

## 🔗 References
- ⬆️ Parent: [[Core Environment Variables]]
- 📚 Module: `Language Basics`
