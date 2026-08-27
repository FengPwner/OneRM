# OneRM

A lightweight, cross-platform command-line utility written in Go. `OneRM` is designed to safely and recursively remove a specific target folder located at the absolute root directory of the system.

## ️ Critical Warning

**This tool is highly destructive.** It permanently deletes the target directory and all of its contents without any backup.
- On Linux/macOS, it targets `/`.
- On Windows, it targets `C:\`.

Please ensure you understand the consequences before running this tool. Always test it in a safe environment first.

## Features

- **Cross-Platform:** Works natively on Linux, macOS, and Windows.
- **Safety Confirmation:** Requires an explicit `Y` input before executing any deletion.
- **Recursive Deletion:** Completely removes the target folder and all nested files/subdirectories.
- **Zero Dependencies:** Built entirely with Go's standard library.

## Installation

### From Source

Ensure you have Go (v1.18+) installed. Clone the repository and build the binary:

```bash
git clone <repository-url>
cd OneRM
CGO_ENABLED=0 go build -ldflags="-s -w" -trimpath -o OneRM .

```
️# CRITICAL WARNING ️

This utility is highly destructive and operates with extreme privileges. 
By running this tool, you acknowledge that it will PERMANENTLY delete 
the target directory and ALL of its contents without any backup.

- On Linux/macOS, the hardcoded target is /.
- On Windows, the hardcoded target is C:\.

There is NO UNDO. Data loss is irreversible. 
Always test in a safe, isolated environment before deploying to production.

Use at your own risk. The author assumes NO liability for any damage, 
data loss, or system instability caused by the execution of this tool.
