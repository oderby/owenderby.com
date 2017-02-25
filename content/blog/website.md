+++
title = "Hosting my Personal Website (for Free!)"
date = "2017-02-25T00:44:37-08:00"
description = "How I built my website"
tags = ["programming"]
slug = "website-setup"
+++

This website came about because I wanted to update my personal website and I had
some time on my hands. After reflection, I decided to build a static site and
host it with GitHub Pages. In retrospect, the process was mostly
straightforward, but there were a few points that weren't obvious to me at the
time. This post details how I built the site, what I tools I ended up using, and
how the site is hosted, for free! Maybe you'll find it helpful!

# Building a Static Site

The first decision I made was that I would use
a [static site](https://en.wikipedia.org/wiki/Static_web_page). I only wanted to
showcase my projects, host my resume, post the occasional blog post, and display
other static content. Any experimentation or fancy interactions I wanted would
most likely be in browser-side Javascript, so a static site would serve my needs
well. Going with a static website also meant I wouldn't need to run my own
server (if I used a third-party hosting service), reducing my potential costs
significantly. Also, because there's no database or server interactions, it
greatly reduces the attack surface, making it much easier to secure the site.

Next, I needed to decide how to build the site. A simple search for "build a
static site" yields an array of tools and even more guides to the "10 best
static site generators". There's also a variety of managed solutions (like
Wordpress, Wix, Squarespace, etc.). Somewhat overwhelming to decide between, to
say the least. I knew I didn't want to do it all from scratch (I didn't want to
style it from scratch or co-opt existing themes), and I was trying to do this as
cheaply as possible, so that ruled out paying for managed solutions. That left
using an open source tool to build it[^1]. But which tool to use?

I wound up going with [Hugo](https://gohugo.io/), because it was well-reviewed,
actively maintained, and written in Go. The documentation is decent (and it was
updated several times just while I was building the site) and the forum is
active, so [getting up and running](https://gohugo.io/overview/quickstart/) was
relatively easy. The tool has quite a few configurations and built-in
integrations (like [Google Analytics](https://gohugo.io/extras/analytics/)
and [Disqus integration](https://gohugo.io/extras/comments/)), providing a lot
of functionality out of the box. There are also a bunch of
pre-built [themes](http://themes.gohugo.io/) to choose from.

One downside I encountered was how cumbersome it is to customize the site in
certain ways. Hugo is great if you everything you want to do fits well within
their existing configurations and exposed functionality (which, to be fair, is
quite expansive). However, if you want to customize anything outside of that,
like the layout of the blog posts page or styling content differently, it
requires diving much deeper into the internals of how Hugo organizes and builds
the site, just to figure out where the bit of logic exists that you want to
change. And then to change it, you need to figure out the local location to
place a file to contain the overriding/customized code. So there's quite a step
function (in the learning curve) between just using the tool and customizing the
result.

Ultimately, I'm glad I went with an existing solution, despite some of it's
rough edges. I didn't explore any other solutions, so I can't say whether Hugo
was the best tool for me or not, but I'm happy with my choice. Although it took
more work to get the site I wanted than I anticipated, I know I got much
farther, quicker with Hugo than I would have on my own.

# Hosting and Serving the Site

Once I had my static site built, I needed to figure out where I'd host it. My
goals for hosting were that it be (1) as cheap as possible and (2) easy to
deploy to. Because it's all static, I could host it on any of the existing file
storage services (e.g. [S3](https://aws.amazon.com/s3/)
or [GCS](https://cloud.google.com/storage/)). However, those would cost money
(likely less than $1 per month, but it's the principle of the thing!), and I had
a hunch I could do it for free. Then I discovered you could host static websites
with [GitHub Pages](https://pages.github.com/). Free and integrated directly
into GitHub (so no extra deploy steps), Pages was clearly the way to go. The
only potential drawback was that my site and code would need to be public. Since
I was using open source tools and all my content was going to be public anyways,
this wasn't a problem for me.

The final piece was figuring out the domain I wanted to host the page at. GitHub
provides free subdomains on the github.io domain for all Pages, but I wanted a
custom domain to truly make the site mine. I had
used [name.com](https://www.name.com/) in the past, so I just used them to
register [owenderby.com](https://owenderby.com) for $8.98 (including Whois
Privacy).

# Gluing Everything Together

With all the major decisions made, it was time to assemble the parts and get the
site up. This was accomplished in the following steps/configurations:

1. Use Hugo to create the initial site: `hugo new site .`
2. [Install a Hugo theme](https://gohugo.io/themes/installing/#installing-a-specific-theme) and
   [set it as default](https://gohugo.io/themes/usage/) for the site.
3. Build the site in the docs directory (this is important for hosting with
   GitHub Pages in step 5): `hugo -d docs`
3. Create a git repository, commit the initial site and push it up to GitHub.
4. Setup GitHub Pages to work for a custom domain. You need to follow
   a
   [few directions](https://help.github.com/articles/using-a-custom-domain-with-github-pages/) to
   do this, including setting up DNS entries and configuring the repository.

One final step I took was setting up TLS (HTTPS) for my site. Although not
strictly required, I generally expect all sites I use these days to be served
over TLS, so my site couldn't be an exception. Sadly, GitHub Pages doesn't
support HTTPS for custom
domains. Fortunately, [Cloudflare](https://www.cloudflare.com/) provides a free
plan for personal websites, which includes their Full SSL solution via shared
certificates, so you can simply sign up for a free account and setup
SSL. Cloudflare has a
good
[blog post](https://blog.cloudflare.com/secure-and-fast-github-pages-with-cloudflare/) covering
how to do this with GitHub Pages, so I just followed the directions there.

And there you go! For the cost of registering a domain name, you have your own
personal static website up and served on your custom domain! And of course,
you could do this all for free if you were satisfied with hosting it on the
GitHub domain.

[^1]: What kind of developer would I be if I didn't build it myself? (And before
    you roll your eyes, at least I didn't build it from scratch!)
