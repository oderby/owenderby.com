+++
tags = [
]
date = "2017-01-09T21:28:52-08:00"
categories = [
]
title = "Exploration of 2016 Presidential Election Results in Pennsylvania"
description = ""
slug = "election_2016_pa_results"
draft = "true"
+++

Since the election, I've spent the past few months exploring the results of the
2016 US Presidential Election in Pennsylvania, one of the "surprising" states to
go to Trump. Obviously, this was motivated out of the basic question "what
happened?" As I dug into the results, I found that, unsurprisingly, there wasn't
any clear, simple answer. Each time I plotted the results, it'd raise new
questions to explore, and so the investigation would continue. Nevertheless, I did discover several interesting
relationships and trends, and after playing around with the data off and on for
the past few months, I wanted to share what I have so far.

While this is yet another post about the 2016 US Presidential Election, this one
differs in two major ways:

1. This post is motivated entirely by data, exploring that data to reveal
   interesting correlations and relationships.
2. This post avoids drawing any conclusions. I am not a political scientist or
   economist, and do not have near enough background to fully interpret the data
   presented here.

What follows is some of the more interesting results I've generated for
Pennsylvania, along with some brief discussion of the methodology, my
interpretation of the results, and questions for further research. I hope you
find it as interesting to read as I did to explore and produce. If anyone wants
to explore the data further themselves, I've made the notebook and data used
available on github. Someday, I'd like to extend this analysis to other states,
but that's limited by time and the availability of county-level data.

All results presented in this post are based on data I collected by manually
scraping and formatting government and public repositories of results and
registration data. Links and methodology can be found on github.

## 2016 Results

To warm up, let's start with the basic results of the 2016 election in
Pennsylvania. This will also allow me to explain the terminology and derived
metrics I use.

First, we note that Donald Trump won the popular vote by 0.72% (44,292 votes). I
will refer to this metric, the difference in votes between the Republican and
Democratic candidate (as a percentage of total votes cast), as the **Vote
Differential** (**VD** for short). Note that VD is positive when the Republican
has more votes, and negative when the Democrat does.

There were a total of 6,115,402 votes cast, with 3.57% (218,228) of votes cast
for neither of the two primary parties. This is the **Independent Vote**
(**IV**).

These votes account for 70.11% of all voters registered in time for the
election. Whenever I talk about voter turnout or other similar values, I use
**Registered Voters** (**RV**) as a substitute for eligible voters or voting age
population. This has many obvious flaws (most important being that not all
eligible voters register, and the difference between registration and eligible
population is likely not uniformly distributed across states), but computing a
more accurate value is far more complex and may not even be possible, due to
relying on data that isn't readily available at a county level for each state.

### County Results
{{< figure src="/img/election_16/pa/results_map_pa_16.png" link="/img/election_16/pa/results_map_pa_16.png" width="100%"
 caption="Fig 1. Map of 2016 election results (VD) by county." >}}

***Methodology*** The above is a map of Pennsylvania. Each county is colored by its
VD, when sampled from a colormap which ranges from dark blue at -1 (Democrat
won), white at 0 (no clear winner), and dark red at +1 (Republican won). The hatches
indicate which counties flipped from 2012 (a different party won this election
than last) and the type of hatch indicates which party won: the dot-hatch means
the democrat won and the cross-hatch means the republican won.

***Discussion*** This map should look familiar, as it's similar to the basic
result map calculated by every media outlet. Glancing quickly at this map, you
might wonder how the election was so close in Pennsylvania, given how many
counties are red and how few are blue. The answer highlights the major drawback
to these sorts of maps - they obscure the affect of the absolute magnitude of
the results across counties. That is, because Pennsylvania decides its electoral
college by popular vote, across the entire state, it's the total number of votes
each candidate receives that matters, not how many counties they won. Many of
the counties which Trump won have a dramatically smaller populations (and thus
votes cast), and thus contribute less towards the popular vote. Indeed, in 2016
Philadelphia county cast 707,631 votes, which comprised more votes cast than the
35 smallest counties combined (that's more than half of Pennsylvania's 67
counties!).

Maps like these also obscure the impact of the IV on the outcome, and given that
the IV for Pennsylvania overall was nearly 5x as large as the VD, it's worth
considering further.

{{< figure src="/img/election_16/pa/results_graph_pa_16.png"
 link="/img/election_16/pa/results_graph_pa_16.png" width="100%"
 caption="Fig 2. Graph of 2016 results, showing independent votes against vote differential." >}}

***Methodology*** Fig 2 plots every county in 2D space, where the x coordinate
is the VD in that county, and the y coordinate is the IV in that county. The
size of each point is proportional to the number of marginal votes that county
yielded for the winning party in that county (so
`abs(VD*county_total_votes)`). The points are colored in the same way as in Fig
1, which is proportional to their VD, and so here color is just a redundant
dimension.

***Discussion*** This is a different way of visualizing the same data as in
Fig 1. Here we can now see a more accurate picture of the results, accounting
for the magnitude of votes each county contributes, as well as how the IV varies
by county. Immediately we can still see how Trump won a majority of the
counties, but it's not immediately clear who has the larger total area, which
would tell us the winner. We can also see that the IV varies greatly by
county. It also seems correlated with `abs(VD)`, but more on that later.

<div class='row'>
    <div class='col-xs-12 col-sm-6'>
        {{< figure src="/img/election_16/pa/iv_16.png" link="/img/election_16/pa/iv_16.png" width="100%"
            caption="Fig 3. IV histogram for 2016." >}}
    </div>
    <div class='col-xs-12 col-sm-6'>
        {{< figure src="/img/election_16/pa/turnout_16.png" link="/img/election_16/pa/turnout_16.png" width="100%"
            caption="Fig 4. Turnout histogram for 2016." >}}
    </div>
</div>

***Methodology*** Figures 3 and 4 are histograms of the IV and turnout,
respectively, by county. Counties are first grouped by the party of the winning
candidate (blue for democrat, red for republican), and then binned.

***Discussion*** I ran this analysis to look for any potential correlation
between which way a county voted and turnout or IV. However, there doesn't
really seem to be - but perhaps a larger sample size would yield different
results?

## Historical Perspective

Looking at this year's results is interesting, but I wanted context. Was this
apparent sweep of counties unprecedented? Afterall, we heard a lot about how
this was the first year in a long time that Pennsylvania went red, so how did
this year actually stack up against previous elections?

To shed some light on these questions, I gathered county-level results for every
presidential election in Pennsylvania since 1996, and registration data
since 2000. Why 1996? Somewhat arbitrary, although it gives us context from the
last two party shifts in the presidency (Bush in '00 and Obama in '08). Also,
the Pennsylvania state website only hosts results since 2000, so I had to get
creative in my google searches. Finally, I hope one day to run this analysis for
other states, and the farther I go back, the more data I'd need to format and
parse to analyze other states. Call me lazy.

### Results Comparisons
{{< figure src="/img/election_16/pa/results_map_years.png"
    link="/img/election_16/pa/results_map_years.png" width="100%"
    caption="Fig 5. Map of election results for years 1996-2016." >}}

***Discussion*** First off, we can immediately see that 2016 doesn't seem *that*
red when compared with 2012. It's definitely redder, but the swing towards the
republican candidate in many counties started long ago. This suggests '08 was an
irregular year in this 16 year trend. Indeed, since 1996 counties have only
flipped democrat in the 2008 election, with the sole exception of Chester county
in '16.

Additionally, we can see the distinct and gradual shift to the right (red) in
the western counties (with the exception of Allegheny, which has Pittsburgh)
since 1996. We'll explore these and other shifts in more detail later on.

{{< figure src="/img/election_16/pa/iv_years.png"
    link="/img/election_16/pa/iv_years.png" width="100%"
    caption="Fig 6. Boxplots of IV distributions for elections 1996-2016." >}}

***Methodology*** Figure 6 presents
the [Tukey boxplots](https://en.wikipedia.org/wiki/Box_plot) for the IV of all
counties by year. Given that Fig. 3 didn't reveal any interesting
variations between county and winning party, I'm ignoring the winning party in
these boxplots.

***Discussion*** Immediately we can see that that votes for independent
candidates increased in '16. Yet, it's no more than was seen in '00, and far
less than in '96. I'd be curious to know if there's any theories on why IV was
so high in '96?

### Year-to-Year Shifts
To get a better feeling for how things have changed over time, looking at the
difference in Vote Differential (the first derivative of VD, or dVD for short)
tells a very interesting story.

{{< figure src="/img/election_16/pa/shifts_map_years.png"
    link="/img/election_16/pa/shifts_map_years.png" width="100%"
    caption="Fig 7. Map of dVD by county for elections 2000-2016." >}}

***Methodology*** Figure 7 presents a map where every county is colored
proportional to its shift in VD (its dVD) from the previous year. To simplify
and better illustrate the general direction and magnitude of shifts, only 7
colors are possible. 3 shades of blue and red each, and white. The shades of red
progress form lightest to darkest, corresponding to 3 ranges of shifts: 3%-10%,
10%-20%, and 20%+. Similarly for blue. White indicates counties which had less
than a 3% shift in either direction.

***Discussion*** This illustrates very clearly the widespread shift to the blue
for Obama in '08, and towards red for Trump in '16. However, perhaps more
surprising is sparsity of unshifting (white) counties. For some reason, I had
expected shifts between elections to be less frequent and more localized. But it
seems like, at least in Pennsylvania over the past 5 elections, the exact
opposite has happened: most counties shifted by >3% in the same direction each
election.


{{< figure src="/img/election_16/pa/dVD_years.png"
    link="/img/election_16/pa/dVD_years.png" width="100%"
    caption="Fig 8. Tukey Boxplots of dVD by election year, 2000-2016" >}}

***Discussion*** Figure 8 summarizes the distribution of VD shifts in each
election. This even more clearly illustrates the point we observed from Figure
7: Aside from 2004, every election since 2000 saw major vote shifts in most
counties, and those shifts often correlate with the winning party. Except for
2012 - when Obama won despite losing significant ground to Mitt Romney.

***Questions***

1. Where there any counties where the VD was less than IV?
2. How does IV correlate with abs(VD)?
3. How much VD did each county contribute to Trump's victory?
4. How does 2016 compare to previous years?
5. How much is dVD accounted for by voters changing sides vs changes in
   turnout/registration?
6. Is 3% VD shift statistically significant?
