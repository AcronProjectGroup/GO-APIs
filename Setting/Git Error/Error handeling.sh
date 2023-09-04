

#  ! [rejected]          main -> main (non-fast-forward)
# error: failed to push some refs to 'github.com:sinalalebakhsh/GO.git'
# hint: Updates were rejected because the tip of your current branch is behind
# hint: its remote counterpart. Integrate the remote changes (e.g.
# hint: 'git pull ...') before pushing again.
# hint: See the 'Note about fast-forwards' in 'git push --help' for details.

# 1
git fetch origin

# 2
git merge origin/your-branch-name
# Save and Out
# Test if OK so OK!


# If you prefer to rebase your local changes on top of the remote changes (which results in a linear history), use the following command:
# 3
git rebase origin/your-branch-name






