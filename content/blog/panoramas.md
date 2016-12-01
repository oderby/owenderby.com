+++
date = "2016-11-29T12:15:19-08:00"
description = "Tips for stitching panoramas and rendering them as a '360° photos' on Facebook"
tags = ["programming", "photography"]
categories = ["programming", "photography"]
title = "Making Panoramas"
slug = "panoramas"
draft = "true"
+++

I recently finished up a month of travel. In addition to seeing great friends
and beautiful places, I used my travels as an opportunity to learn how to use a
DSLR. I mean *really learn*, beyond just the point-and-shoot basics we all
know. So I borrowed an old Nikon D40 from a friend and sort of just dove
in. Turns out, it's not very difficult to go from n00b to novice, and it's even
easier to take a lot of photos...but that's a story for another day.

One of the challenges I (unwittingly) set for myself was creating panoramas. It
seems straightforward in theory - take a series of overlapping photos and then
combine them using software. Afterall, most modern smart phones have panorama
modes built into their camera apps. However, as I soon learned, it can very
easily turn into a deep, interesting ~~trap~~rabbit hole of techniques and
details, both on the photo-taking and post-processing side.

This post is a summary of my findings from figuring out how to create panoramas
from individual photos, and then uploading them to Facebook
as "Photo Spheres"[^6]. This wasn't as simple as I expected it to be, so I wanted to
record my insights and process, for my future self and anyone else who is silly
enough to try and do this.

One last aside - I work predominantly on Ubuntu with FOSS, so this is written
assuming you're using a similar environment. This is definitely possible using
Photoshop (heck, it's probably much easier), but YMMV.

## Making the Panorama

There's effectively two components to my process:

<ol type="A">
<li>Produce the raw panorama by stitching together a series of pictures.</li>
<li>Set the panorama to have the correct field of view (FOV) and image size and
   then add appropriate metadata so it's recognized as a Photo Sphere.</li>
</ol>

This section focuses on part A.

### Taking the Photos

Since I was just starting out with a borrowed camera, I wasn't very
sophisticated in my setup for taking the panorama shots. I shot hand-held, and
used the horizon as my reference as I swept my camera across the scene, taking
pictures periodically. Of course, there are
some [very fancy rigs](http://wiki.panotools.org/Heads) out there to buy for
your tripod, and there are even some clever techniques for better hand-held
shots (google "philopod" if you're curious).

Nevertheless, I found I was able to produce decent panoramas, once I learned how
to keep the exposure and white balance the same in every picture.

> ***Insight #1*** – Consider using
> the [AutoExposure-Lock](https://photographylife.com/nikon-ae-l-af-l-button)
> feature of your camera when taking a photo series for a panorama. It makes the
> resulting panorama look more natural and contiguous, at the cost of some
> places being incorrectly exposed.

I didn't have many problems with parallax, since I was taking mostly wide-angle
landscape panoramas. But I now know that I can reduce parallax effects by
rotating the camera about
its ["no-parallax point"](http://wiki.panotools.org/No-parallax_point).

### Stitching the Photos

For stitching, the best open source tool I found[^1]
is [Hugin](http://hugin.sourceforge.net/). Hugin is actually a front end GUI for
the [Panorama Tools](http://wiki.panotools.org/Panorama_tools).

Now, Hugin is an amazing tool and can be used to do a tremendous amount of
things. And so, here is where the hole really starts. Between trying to figure
out how to use the tool and being distracted by learning what people do with the
tool, I spent way too long here. And every time I revisit this, I seem to lose
more time[^2]. But, even I have limits to how deep down the rabbit hole I'll
go. So, here I'll be focusing on the practical findings, and leaving some
questions unanswered...because...I just need to move on :dizzy_face:

So, without further adieu, here's the CliffsNotes[^4] for creating panoramas in
Hugin:

1. Load you photo series.
2. Switch to 'Advanced' interface.
3. On the 'Photos' tab:
  1. Select 'Hugin\'s CPFind + Celeste' option for Feature Matching. Press
     'Create control points' button.
  2. Select 'Positions, View and Barrel' for Geometric.
  3. Run the optimizer (Edit > Optimize)
4. Switch back to 'Simple' interface.
5. Go to 'Preview' tab and review panorama. You're looking for a few things:
  1. Are there any redundant photos? (e.g. photos arranged \[1\]\[2\]\[3\] such
     that 1 and 3 have enough overlap that 2 isn't needed).
  2. Are there any moving objects (cars, people) in the image? If so, consider
     using a [mask](http://wiki.panotools.org/Hugin_Mask_tab) to hide it.
  3. Does the panorama look good (well aligned and not distorted)? If not, you
     may need to manually
     add/remove
     [control points](http://wiki.panotools.org/Hugin_Control_Points_tab).
6. Go to the 'Projection' tab and press the 'Fit' button. You may want to try
   different projections to see if any work better. For wide-angle panoramas of
   nature, 'cylindrical' seems best.
7. Go to the 'Move/Drag' tab and press 'Straighten' button. This should make a
   very subtle change to the panorama - if instead it turns your panorama into
   an 'S', Hugin probably has your lens information or focal length wrong[^5].

***midway point image***

Okay, so now at this point you have a choice, depending on what sort of panorama
you want to produce. Typically, you'll want to produce a Photo Sphere, which is a
much cooler way to show off your panorama. In that case, proceed to the
next [section]({{< relref "#make-sphere" >}}) now.

However, sometimes you just want a flat, cropped panorama. I typically opt for
this when the estimated horizontal FOV[^8] is ≤160° or I'm just stitching together a
couple photos (not really a panorama). To do this, continue:

1. Switch to 'Advanced' interface.
2. Go to 'Stitching' tab and...
  1. Press 'Calculate Optimal Size'
  2. Press 'Fit Cropt to Images'
  3. Ensure 'Exposure corrected, low dynamic range' is the only option selected.
  4. Change the format to 'JPEG'
  5. Press 'Stitch!' button and choose place to save hugin project file, then
     name and place to save resulting panorama.

That's it - you should have a nice, clear panorama to post as you like!

***flat, cropped panorama image***

## Making a Photo Sphere[^12] {#make-sphere}

Now, the panorama is basically done, you just need to massage it into a format
that Facebook will accept and recognize as a 360° image[^9]. To do so, several
important steps remain.

### Image Shape

First, you need to make your panorama fit into one of seven predefined sizes
(determined by horizontal and vertical FOV):

1. Go back to the 'Projection' tab and note the estimated horizontal and
   vertical FOV values. Then find the row in the following table that best
   matches those values. Enter the FOV values from the table into Hugin. The
   Image Size and Command you'll use next.

    | H FOV | V FOV | Image Size (px) | Command |
    | --- | --- | --- |  --- |
    | 120° | 86° | 6000x5343 | {{< copytext >}}exiftool -overwrite_original -FullPanoWidthPixels=18000 -FullPanoHeightPixels=9000 -CroppedAreaLeftPixels=6000 -CroppedAreaTopPixels=2350 -CroppedAreaImageWidthPixels=6000 -CroppedAreaImageHeightPixels=4300 -ProjectionType=cylindrical path/to/image.jpg{{< /copytext >}} |
    | 150° | 86° | 6000x4274 | {{< copytext >}}exiftool -overwrite_original -FullPanoWidthPixels=14400 -FullPanoHeightPixels=7200 -CroppedAreaLeftPixels=4200 -CroppedAreaTopPixels=1880 -CroppedAreaImageWidthPixels=6000 -CroppedAreaImageHeightPixels=3440 -ProjectionType=cylindrical path/to/image.jpg{{< /copytext >}} |
    | 180° | 86° | 6000x3562 | {{< copytext >}}exiftool -overwrite_original -FullPanoWidthPixels=12000 -FullPanoHeightPixels=6000 -CroppedAreaLeftPixels=3000 -CroppedAreaTopPixels=1567 -CroppedAreaImageWidthPixels=6000 -CroppedAreaImageHeightPixels=2867 -ProjectionType=cylindrical path/to/image.jpg{{< /copytext >}} |
    | 240° | 86° | 6000x2671 | {{< copytext >}}exiftool -overwrite_original -FullPanoWidthPixels=9000 -FullPanoHeightPixels=4500 -CroppedAreaLeftPixels=1500 -CroppedAreaTopPixels=1175 -CroppedAreaImageWidthPixels=6000 -CroppedAreaImageHeightPixels=2150 -ProjectionType=cylindrical path/to/image.jpg{{< /copytext >}} |
    | 300° | 86° | 6000x2137 | {{< copytext >}}exiftool -overwrite_original -FullPanoWidthPixels=7200 -FullPanoHeightPixels=3600 -CroppedAreaLeftPixels=600 -CroppedAreaTopPixels=940 -CroppedAreaImageWidthPixels=6000 -CroppedAreaImageHeightPixels=1720 -ProjectionType=cylindrical path/to/image.jpg{{< /copytext >}} |
    | 360° | 86° | 6000x1781 | {{< copytext >}}exiftool -overwrite_original -FullPanoWidthPixels=6000 -FullPanoHeightPixels=3000 -CroppedAreaLeftPixels=0 -CroppedAreaTopPixels=783 -CroppedAreaImageWidthPixels=6000 -CroppedAreaImageHeightPixels=1433 -ProjectionType=cylindrical path/to/image.jpg{{< /copytext >}} |
    | 360° | 65° | 6000x1217 | {{< copytext >}}exiftool -overwrite_original -FullPanoWidthPixels=6000 -FullPanoHeightPixels=3000 -CroppedAreaLeftPixels=0 -CroppedAreaTopPixels=958 -CroppedAreaImageWidthPixels=6000 -CroppedAreaImageHeightPixels=1083 -ProjectionType=cylindrical path/to/image.jpg{{< /copytext >}} |

2. Go to the 'Crop' tab and click on the button next to the crop dimensions to
   'reset crop to maximal possible area'. The result should be your panorama
   embedded in an overall image which covers the desired FOV.
3. Switch back to 'Advanced' interface.
4. Go to 'Stitcher' tab and...
  1. Change the Canvas Width to be 6000. The Height value should match the Image
     Size given in the above table. Leave all other values for FOV and Crop
     alone.
  2. Ensure 'Exposure corrected, low dynamic range' is the only option selected.
  3. Change the format to 'JPEG'
  4. Press 'Stitch!' button and choose place to save hugin project file, then
     name and place to save resulting panorama.

This should result in something that looks like the following:

***Flattened 360 panorama***

### Image Metadata

Now, you need to set the metadata correctly on your image so that it *declares*
itself as a Photo Sphere. This can be done with
the [exiftool](http://owl.phy.queensu.ca/~phil/exiftool/) command line
tool. You'll notice in the table above I've provided the appropriate commands to
run for each FOV. Copy the appropriate command and run it in your terminal[^13]. Be
sure to replace `path/to/image.jpg` with the appropriate path to your image.

### Uploading to Facebook

Finally! You're ready to upload to facebook...or are you? Well, you are, but you
can only upload to the web (desktop) version of Facebook, and you have to post
it directly to your timeline (you can't add it to an album directly[^14]), and you
can only upload one panorama at a time.

## Future Directions
That's it! Hope it was useful, or at least insightful...

If I have time, I'd like to investigate the following areas:

1. Investigate how to use [bracketing](https://en.wikipedia.org/wiki/Bracketing)
   to get a better overall exposure for the panoramas.
2. What is lens [vignette](https://en.wikipedia.org/wiki/Vignetting), and do I
   need to correct/remove it?
2. I find the max 6000 pixel photo size imposed by facebook to be limiting. So
   I'd like to investigate ways to provide my own 360° photo viewer, using
   three.js. [This](https://threejs.org/examples/?q=panorama#webgl_panorama_equirectangular) looks
   promising.

[^1]: Alright, the only one I found that I actually tried. But after further searching, it seems to actually be the best one out there.

[^2]: I just fell, again, into the deep ~~trap~~rabbit hole again figuring out where to point to for a good summary document here...

[^4]: Yes, that is the [correct spelling](https://en.wikipedia.org/wiki/CliffsNotes)...

[^5]: This is an area that, despite several attempts, I haven't fully figured out yet. The focal length and other important lens information is always present in the photo EXIF data, but for some reason Hugin typically inputs some different values for focal length and focal length multiplier (crop factor), and I haven't figured out why yet. Typically, just leaving those alone seems fine, but if my panorama isn't aligning properly, I'll try putting in the exact values from the EXIF data, and that does the trick.

[^6]: Technically, a panorama using a spherical projection for viewing. Google calls these "Photo Spheres", and that's the term I'll use. However, Facebook calls them ["360° photos"](https://facebook360.fb.com/360-photos/). I dislike this term because they're not always 360° (in particular, for most panoramas you want to take with your DSLR.

[^8]: This is another area I don't fully understand - how does Hugin estimate the field of view for the resulting panorama?? I'm pretty sure the math is outlined [here](http://wiki.panotools.org/DSLR_spherical_resolution), but I think that article is a bit muddled...

[^9]: This was an especially infuriating endeavor, as Facebook has [detailed documentation](https://facebook360.fb.com/editing-360-photos-injecting-metadata/) on how to do it, but it's focused mostly on producing 360° Photo Spheres using Photoshop. It took a lot of trial and error to figure out how to make it work for partial Photo Spheres.

[^12]: For uploading to Facebook. Although the result should work with any other Photo Sphere renderer, I haven't tested it.

[^13]: For the curious, this command specifies how to place our partial panorama inside a full spherical panorama (360° x 180°). See [here](https://developers.google.com/streetview/spherical-metadata#metadata_properties) for further details and a nice illustration. The astute will observe that this is redundant with the work we already did to fit our original panorama into one of the predefined FOV Facebook specifies...and all I can say is "Yes. Don't ask."

[^14]: Instead, you can move a photo from your timeline to an album. Just open the image, then go to Options > 'Move to Other Album'.
