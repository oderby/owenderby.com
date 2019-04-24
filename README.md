# Owen Derby's Website

This repository is my personal website. I use [Hugo](https://gohugo.io/) to
compile the source into a static website in the `docs/` directory. The latest
website can always be found at [https://owenderby.com](https://owenderby.com).

I followed these tutorials ([1](http://www.moxie.io/post/static-websites-with-hugo-on-google-cloud-storage/) [2](https://gohugo.io/overview/quickstart/)) to get up and running with Hugo.

# Building

Copy [.env.sample](.env.sample) to .env and fill in phone number.

Run `make serve`, preview the website at http://localhost:1313 and view the [private resume pdf](resume.pdf).

Run `make site` to generate the site to the docs/ folder, then commit and push changes to `master` branch.

# Scratchspace

## Possible Themes:

* http://themes.gohugo.io/theme/beautifulhugo/
* http://themes.gohugo.io/theme/hugo-sustain
* http://themes.gohugo.io/theme/hugo-icarus/
* http://themes.gohugo.io/theme/internet-weblog/

## Useful tools (for future expansions)

* https://github.com/exif-js/exif-js/tree/master/example
* https://github.com/bep/bepsays.com
* https://github.com/Kris-B/nanoGALLERY
* https://github.com/mwhite/resume
* https://docs.travis-ci.com/user/deployment/gcs/
* http://paperjs.org/about/
* https://imulus.github.io/retinajs/

## Future Ideas/Work:

* Provide contact method
  - Something like [this](https://github.com/dwyl/html-form-send-email-via-google-script-without-server)?
* Printable resume
* Improve google analytics
* Support 360° photospheres
  * https://github.com/mrdoob/three.js/blob/dev/examples/webgl_panorama_equirectangular.html
  * https://chris.orr.me.uk/photo-sphere-embedding-viewer-tiny-planet/
  * https://github.com/kennydude/photosphere
* cleaner format for blogs/ page
* Favicon
* Uglify/Minify code (mostly as an exercise, not because I care about the load
  performance)

# Credits

This repository and website pulls ideas and source code from the following
places:

* [Steve Francia's Website](http://spf13.com/)
  and [Source Code](https://github.com/spf13/spf13.com)
* [Bjørn Erik Pedersen's website repository](https://github.com/bep/bepsays.com)
* [Nurlan Su's Sustain theme for Hugo](https://github.com/sumaxime/hugo-sustain/)

# License

The intent behind the licensing for this repository is that all source code for
site layout and customizations (css, js, raw html, etc.) is MIT licensed, while
all original textual and image content (the "content" for the website) is
Copyrighted by Owen Derby (or their respective copyright holders) and no rights
are granted, unless explicitly indicated.

That roughly translates into the following files and directories (including
their contents) are Copyright Owen Derby (all rights reserved) or their
respective copyright holders and licenses:

  * content/
  * docs/
  * static/img/
  * static/favicon.ico
  * static/apple-touch-icon.png
  * themes/

All other directories and files are MIT Licensed, unless clearly
designated otherwise.
