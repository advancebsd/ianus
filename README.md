# Iᴀɴᴜs
![logo](/ianus.png "Iᴀɴᴜs")

Named after the two-faced Roman god Janus (in its original Latin spelling), Iᴀɴᴜs is the attempt of implementing the concept of a dual-protocol server for HTTP and [Gemini](https://en.wikipedia.org/wiki/Gemini_(protocol)). This can be useful for people who regularly publish material both on the Web and in Gemini space (so-called *bi-posting*).

To make this as feasible as possible, the application will serve both HTML and Gemtext content that is automatically generated from a common intermediary source. The latter is read directly from a git repository (per vhost) to support a convenient site development process.

Once the basic functionality is considered stable, several advanced features may be added.
