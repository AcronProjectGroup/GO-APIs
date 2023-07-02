# The standard debugger for Go applications is called Delve.

# Project:
#     https://github.com/go-delve/delve


go install github.com/go-delve/delve/cmd/dlv@latest
# Result:
    # go: downloading github.com/go-delve/delve v1.21.0
    # ^Bgo: downloading github.com/sirupsen/logrus v1.6.0
    # go: downloading github.com/spf13/cobra v1.1.3
    # go: downloading github.com/mattn/go-isatty v0.0.3
    # go: downloading github.com/cosiner/argv v0.1.0
    # go: downloading github.com/derekparker/trie v0.0.0-20221213183930-4c74548207f4
    # go: downloading github.com/go-delve/liner v1.2.3-0.20220127212407-d32d89dd2a5d
    # go: downloading github.com/google/go-dap v0.9.1
    # go: downloading go.starlark.net v0.0.0-20220816155156-cfacd8902214
    # go: downloading github.com/hashicorp/golang-lru v0.5.4
    # go: downloading golang.org/x/arch v0.0.0-20190927153633-4e8777c89be4
    # go: downloading golang.org/x/sys v0.0.0-20220908164124-27713097b956
    # go: downloading github.com/mattn/go-runewidth v0.0.13
    # go: downloading github.com/cilium/ebpf v0.7.0
    # go: downloading github.com/cpuguy83/go-md2man/v2 v2.0.0
    # go: downloading github.com/rivo/uniseg v0.2.0
    # go: downloading github.com/russross/blackfriday/v2 v2.0.1
    # go: downloading github.com/shurcooL/sanitized_anchor_name v1.0.0


dlv version
# Result:
    # Delve Debugger
    # Version: 1.21.0
    # Build: $Id: fec0d226b2c2cce1567d5f59169660cf61dc1efe $



# If you receive an error that the dlv command cannot be found
~/go/bin/dlv
# If the package has been installed correctly, you will see the output similar
# Delve Debugger
# Version: 1.7.1
# Build: $Id: 3bde2354aafb5a4043fd59838842c4cd4a8b6f0b $



