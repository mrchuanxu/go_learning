#### git-am
Apply a series of patches from a mailbox<br>
Splits mail messages in a mailbox into commit log message, authorship information and patches, and applies them to the current branch.<br>


```
git am [--signoff] [--keep] [--[no-]keep-cr] [--[no-]utf8]
	 [--[no-]3way] [--interactive] [--committer-date-is-author-date]
	 [--ignore-date] [--ignore-space-change | --ignore-whitespace]
	 [--whitespace=<option>] [-C<n>] [-p<n>] [--directory=<dir>]
	 [--exclude=<path>] [--include=<path>] [--reject] [-q | --quiet]
	 [--[no-]scissors] [-S[<keyid>]] [--patch-format=<format>]
	 [(<mbox> | <Maildir>)…​]
git am (--continue | --skip | --abort | --quit | --show-current-patch)
```


