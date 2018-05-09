# atmysqldump

This relatively small [Go language](https://golang.org/) based application was written out of a bit of frustration that resulted from a repeatedly realized need to do MySQL backups via cron (or something like it) from within a Docker container as a non-root user.  While this sounds like a relatively simple objective, at the time I wrote the initial version of this application, I did not get that impression.

In fact, I would say it was doable but I came to the realization that I could spend a couple hours writing a small Go-based app in less time than it would take to make cron or one of the few alternatives run as non-root.  Perhaps I missed something but I did have fun writing this.  If nothing else it served as a nice reminder of how much I like this language.

