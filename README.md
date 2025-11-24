# "Of course they'll be stuck."

You already know this story — the new dev joins, the README is outdated, and your senior loses half a day fixing Node versions and missing .envs.

`devcheck` is the one-command sanity check that fixes all that.

**Current Status:** v0.1.0 (Public Beta)

---

## ⚡ Quick Start

### Option 1: The "Zero-Dependency" Way (Recommended for Teams)
No `pip`, `npm`, or `go` required.

1.  **Download** the binary for your OS from the [Releases Page](../../releases).
2.  **Commit it** to your project (e.g., inside a `./tools` folder).
3.  **Run it:**
    ```bash
    ./tools/devcheck
    ```

### Option 2: The Go Developer Way
If you already have Go installed:
```bash
go install [github.com/ishwar170695/devcheck-idea@latest](https://github.com/ishwar170695/devcheck-idea@latest)
devcheck
````

-----

## 🛠 Configuration

Create a `devcheck.toml` file in the root of your project. Think of this like `ruff.toml`, but for your environment.

```toml
# devcheck.toml

[versions]
# Checks output of "node -v", "python --version", etc.
node = "18.x"
python = "3.10"
docker = "24."

[env]
# Checks if these variables exist in the current session
keys = ["DATABASE_URL", "API_KEY", "STRIPE_SECRET"]

[connections]
# Runs any shell command. Exit code 0 = PASS.
# Useful for checking DB connectivity or Docker health.
db_ping = "ping -c 1 $DATABASE_URL"
```

**The Result:**

```bash
$ devcheck

[NODE]   v18.1.0 ... PASSED
[ENV]    API_KEY ... FAILED
[DB]     Ping...   ... FAILED

[devcheck FAILED] Missing .env key detected.
```

-----

## 🧠 Why? (The "Jugaad" is the Bug)

Your onboarding isn't "a process." It's a collection of broken "jugaad" that everyone is too busy to fix.

### 1\. The Messy README

The doc that "was outdated the day after the last intern left." Every setup guide becomes fiction in a month.

> “Usually it’s the messy readme or the dev losing 2-3hrs or both.”
> \~ Senior Dev, r/developersIndia

### 2\. The Senior Hand-Holding

The "fix" where you **"assign one of your senior engineers for a couple of weeks"** just to get one person's environment running.

### 3\. The Docker 'Fix'

The "gold standard" that still leads to new headaches.

> “I tried this but guess what now I spend hours trying to setup docker on their machine”
> \~ Backend Dev, r/developersIndia

### 4\. The Blame Culture

The "works on my machine" reply that kills morale. `devcheck` is the **Objective Referee** that proves it's a missing config, not a "bad developer."

-----

## 🤝 Contributing & Feedback

This is a "dumb" tool by design. It doesn't try to be AI. It just tries to be the "Step 0" that runs before everything else.

If you find a bug or have a feature request, please [Open an Issue](https://github.com/ishwar170695/devcheck-idea/issues) or [DM me on Reddit](https://www.reddit.com/user/byte4justice).

```
```
