+++
tags = [
]
date = "2017-02-03T21:28:52-08:00"
categories = [
]
title = "Exploration of 2016 Presidential Election Results in Pennsylvania"
foo = "An examination of 2016 Election Results in Pennsylvania"
description = ""
slug = "election_2016_pa_results"
draft = "true"
+++
This past election has been on my mind a lot lately. When it comes to society
and politics, I've learned that the issues are often complex and obfuscated. We
can only find clarity by digging deeper and elucidating the nuances of the
issues. Unlike legislation and policy, election results are far more objective,
lending themselves to analysis. And so, shortly after the election, I dove into
the data to seek clarity about the dynamics of this election.

I decided to start with Pennsylvania state, as a smaller, more tractable case
study. I examined registration and results data for this and previous
elections. My analysis has yielded two general insights about this
election. First, this election is remarkable, not because Trump won, but because
of how he won. Second, that a Republican won Pennsylvania this year is only
remarkable when compared to Obama's win in 2008.

Trump's win in Pennsylvania is certainly remarkable. It is the first time a
Republican candidate won the state in 28 years. Moreover, while the race with
his opponent (Clinton) was close, he outperformed the previous Republican
candidate (Romney) by a wider margin than Obama did over his predecessor (Kerry)
in 2008. That is, **Trump swung more votes in 2016 than Obama did in 2008**.

Which brings me to my second point. 2008 was itself a remarkable election, no
doubt. And when we compare 2016 to 2008, the differences seem stunning. Yet,
when we step back and put things into the context of the surrounding elections,
we see that **2008 was the anomaly, not 2016**. In fact, Republican candidates
have been steadily gaining ground in Pennsylvania since 1996 (the earliest I was
able to get data for). 2008 is the only year where that trend was reversed and
Obama gained tremendous ground while the Republicans lost so much. However, in
2012 Obama lost much of that ground and the Republicans were back on track.

In this post, I examine the election data in detail to illustrate these points
and try to explain


<!---
I'm an engineer. As an engineer, I like it when things are clear and logical. In
fact, I consider it part of my job to make things as clear as possible.

> Trump won the election, but Clinton won the popular vote.

The above fact is anything but clear and logical, and yet it's
indisputable. How? What happened?

In my quest to answer these questions, I came across another, lesser known, but
equally confounding fact.

> In Pennsylvania, Trump won his counties by larger margins than Obama did in
> 2008.
!--->

Since the election, I've spent the past few months exploring the results of the
2016 US Presidential Election in Pennsylvania, one of the "surprising" states to
go to Trump. Obviously, this was motivated out of the basic question "what
happened?" As I dug into the results, I found that, unsurprisingly, there wasn't
any clear, simple answer. Each time I plotted the results, it'd raise new
questions to explore, and so the investigation would continue. Nevertheless, I
did discover several interesting relationships and trends, and after playing
around with the data off and on for the past few months, I wanted to share what
I have so far.

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

## So What Happened?

Let's start with the basics:

<!---
Such a broad and backwards-looking question, but it's the one I've heard
repeated more times than I can count since November 9, and just as often within
my own head as without. It's almost pointless to ask, and yet it's as good a
starting point as any.

Let's begin by taking the question literally.
!--->
<!---
For simplicity and to keep the
analysis tractable, I started by analyzing a single state: Pennsylvania. And so,
we begin with a basic summary of the 2016 election in Pennsylvania. This will
also allow me to explain the terminology and derived metrics I use.
!--->

* Donald Trump won the popular vote by 0.72% (44,292 votes). I will refer to
  this metric, the difference in votes between the Republican and Democratic
  candidate (as a percentage of total votes cast), as the **Vote Differential**
  (**VD** for short). Thus, for the entire state, the VD was +0.0072. Note that
  VD is positive when the Republican has more votes, and negative when the
  Democrat does.
* That 0.72% difference corresponds to 44,292 **Marginal Votes** (**MV**) for
  Trump, at the state level. MV[^4] is just the difference in votes for the
  Republican and Democrat candidates. Note that `VD=MV/Total Votes`.
* In contrast, Barack Obama won the election in 2012 by 5.45% (so in 2012 the VD
  for the state was -0.0545). I will often want to compare changes in VD and
  other metrics between elections. When I do this, I will compute the change[^2] as
  `new_value-old_value` and prefix the metric with a 'd' (for 'delta' or
  'difference'). Thus, the **dVD** for the state for 2012-2016 is
  `0.0072-(-0.0545)=0.0617`.
* There were a total of 6,115,402 votes cast, with 3.57% (218,228) of votes cast
  for neither of the two primary parties. This is the **Independent Vote**
  (**IV**).
* The total votes cast account for 70.11% of all **Registered Voters** (**RV**)
  for the election. For simplicity, I will use this (votes cast over registered
  voters) as a measure of voter **Turnout**[^1]. There were 3,301,543 registered
  Republicans and 4,217,094 registered Democrats, giving a **Registration
  Differential** (**RD**, the analog of VD for registration statistics) of
  -0.105.


### But, How Did Trump Win?

From above, it's apparent that Trump won, but the state-wide numbers don't
really help us understand how Trump won[^3]. To get a better understanding of
what was going on, we need to dive deeper and look at county-level results. We
start by examining the results for each county on a map. Perhaps we'll find some
correlation into how voters voted and where they lived.

{{< figure src="/img/election_16/pa/results_map_pa_16.png"
 link="/img/election_16/pa/results_map_pa_16.png" width="100%"
 caption="Fig 1. Map of 2016 election results (VD) by county. Cross-hatches indicate counties that Obama won in '12 which 'flipped' to Trump in '16. Dot-hatches indicate the reverse (Romney->Clinton)." >}}

<!--- ***Methodology*** In the above map of Pennsylvania, each county is colored
by its VD, when sampled from a colormap which ranges from dark blue at -1
(Democrat won), white at 0 (no clear winner), and dark red at +1 (Republican
won). The hatches indicate which counties flipped from 2012 (a different party
won this election than last) and the type of hatch indicates which party won:
the dot-hatch means the democrat won and the cross-hatch means the republican
won.  !--->

This map should look familiar, as it's similar to the basic result map
calculated by every media outlet. Glancing quickly at this map, you might wonder
how the election was so close in Pennsylvania, given how many counties are red
and how few are blue (Trump won 56 counties to Clinton's 11). The answer
highlights the major drawback to these sorts of maps - they obscure the effect
of the absolute magnitude of the results across counties. That is, because
Pennsylvania decides its electoral college by popular vote, across the entire
state, it's the total number of votes each candidate receives that matters, not
how many counties they won. Many of the counties which Trump won have a
dramatically smaller populations (and thus votes cast), and thus contribute less
towards the popular vote. Indeed, in 2016 Philadelphia county cast 707,631
votes, which comprised more votes cast than the 35 smallest counties combined
(that's more than half of Pennsylvania's 67 counties!).

***Talk about location of 4-5 largest cities in PA?***

Figure 2 attempts to better illustrate the impact different counties have on
each candidate's overall outcome, by taking into account the number of votes each
county generated for each candidate.

<!---
It's important to note that both Figures 2 and 3 are dealing with *marginal
votes*, which is the number of votes in a county that the winning candidate
received beyond what the losing candidate lost. This is related to the VD metric
defined above, except it's an absolute measure, in terms on number of votes (VD
is normalized by total number of votes). This means, when comparing marginal
votes between two elections, the marginal vote in a county can change due to one
candidate receiving more votes or the other receiving fewer votes can both cause
shifts in the same direction, or both. For example, in 2012 in Erie county,
Romney won 49,025 votes against Obama's 68,036, giving Obama 19,011 marginal
votes. In 2016, Trump won 60,069 votes against Clinton's 58,112, yielding 1,957
marginal votes for Trump. And so the change in marginal votes for Erie county
between 2012 and 2016 was 20,968 votes, due to Clinton receiving 9,924 fewer
votes than Obama and Trump gaining 11,044 votes.
!--->
<!---
Maps like Figure 1 also obscure the impact of the IV on the outcome, and given that
the IV for Pennsylvania overall was nearly 5x as large as the VD, it's worth
considering further.

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
!--->

{{< figure src="/img/election_16/pa/results_marginal_votes_barh_pa_16.png"
    link="/img/election_16/pa/results_marginal_votes_barh_pa_16.png" width="100%"
    caption="Fig 2. Stacked bar chart of county marginal votes by winning party. Each bar-section is a county, with width=MV and color=VD (same color as in Fig 1). Sections are ordered by VD." >}}

<!--
***Methodology*** Figure 2 presents a visualization of the marginal vote for
every county as a stacked bar chart. The width of each section of each bar
corresponds to the number of marginal votes that party's candidate won in a
county, while the color is the same as in Figure 1 (colormap of VD for each
county). For ease of visualization, the bar sections are ordered by VD (which is
why both bars go from dark to light, corresponding to high absolute VD to lower
VD).
!-->

First, note that once we take the absolute impact of each county into account,
the race appears much closer (as seen by the closeness of the total size of the
two bars). Second, also note how clearly this figure illustrates the outsized
influence Philadelphia county has, as discussed above. Finally, notice how the
number of marginal votes for republicans don't seem to be correlated strongly
with VD (if they were, we'd expect to see each section get progressively smaller
from left to right, as the VD and marginal vote yield decreased simultaneously).

But to really understand how Trump won, we need to understand the delta - how
things changed since the last presidential election in 2012.

{{< figure src="/img/election_16/pa/results_marginal_votes_barh_pa_12.png"
    link="/img/election_16/pa/results_marginal_votes_barh_pa_12.png" width="100%"
    caption="Fig 3. Stacked bar chart of marginal votes by winning party." >}}

Now, here's where things start to get interesting. Comparing Figure 3 with
Figure 2 tells a striking story. While Clinton garnered as nearly as many
marginal votes as Obama did in '12 (she was short by 3,728 votes), Trump earned
75% more than Romney (Trump gained 353,065 additional votes). And Trump won
these additional votes from nearly the same counties as Romney (only 4 counties
flipped from '12 to '16, as Figure 1 showed).

So where did these marginal votes for Trump come from? To better understand
this, I examine various theories. For each theory, I look to see what the data
shows and if it supports the theory, how many of Trump's additional 353,065
marginal votes it could explain. All comparisons continue to be made against
2012 results.

### Theory #1: Republicans had Higher Turnout

Probably the most credible theory is that republican voters turned up in higher
numbers than democrats in many counties. As Figure 4 shows, in counties Trump
won, the turnout was about 3% higher on average (and as much as 8% higher in
some counties) than those Clinton won. 3% may not seem like a lot, but remember,
Trump won a lot of counties. So, if we pessimistically assume all these
additional votes voted for Trump, that would represent 220,794 marginal votes
(about 57% of the marginal votes Trump gained over Romney).

{{% fig src="/img/election_16/pa/dturnout_16.png" %}} Fig 4. Map (left)
and [Tukey boxplots](https://en.wikipedia.org/wiki/Box_plot) (right) of
distribution of d(turnout) by winning party in 2016." {{% /fig %}}


But republicans could have turned up in a different way - by increasing the
relative number of registered Republican voters since 2012. Interestingly, this
can happen through three different mechanisms - new voters can register
republican, registered democrats could have switched parties, or registered
democrats could have been un-registered. More on this later ***INSERT LINK TO
SECTION***

{{% fig src="/img/election_16/pa/dRD_16.png" %}} Fig 5. Map (left)
and [Tukey boxplots](https://en.wikipedia.org/wiki/Box_plot) (right) of
distribution of dRD by winning party in 2016. {{% /fig %}}

As Figure 5 shows, relative registration increases were even more significant
between counties Trump won and counties Clinton won. They differed by about 6%
on average, and represent (again, taking a pessimistic view that all these extra
registered Republicans turned out to vote for Trump) as much as 227,709 marginal
votes for Trump, about 60% of the marginal votes he gained over Romney.

And while that definitely could potentially explain a majority of Trump's win,
it's unlikely that it explains every vote, and assumes a very perfect situation,
where every possible vote due to higher turnout or registrations voted for him -
when in all likelihood these extra voters turned up in the same proportion as
the rest of voters and voted closer along party lines. Using these more
realistic assumptions, we'd attribute just 26,528 marginal votes (7%) from
higher turnout and 160,250 marginal votes (42%) from shifting registrations to
Trump. Note there is also likely some overlap in the impact of these phenomenon
on these numbers (increased turnout is related to increased net republican
registrations).

**Conclusion:** While increased republican turnout and registrations across the
state likely account for a large portion of Trumps extra marginal votes (close
to 50%, using the reasonable estimates), these phenomenon don't explain them
all.

### Theory #2: Independent Candidates and Disillusioned Democrats

Perhaps the other common theory I've heard is that many democrats were
disillusioned with Clinton (including so-caled "Bernie Bros", amongst
others). The theory holds that these voters opted to vote for an independent
candidate instead of along party lines, or simply didn't turn out to vote at
all, and that's why Trump was able to win a historically 'blue' state. Can we
see if this is true? As with all of these questions, it's hard to ascertain the
truth, and nearly impossible to get a definitive answer with the data I'm
using. Nevertheless, there are a few things we can look at.

First, we consider the possibility of a lower turnout by democrats. However,
Figure 4 shows that turnout increased compared to 2012 in nearly all
counties. And while it's possible that much of that turnout was republicans
voters, it's unlikely it was exclusively republicans. In any event, we've
already covered what impact that had on Trump's victory in the previous section
(whether more Trump voters turned out or fewer would-be Clinton voters turned
out, the result on marginal votes is the same), and we'll cover how it could
have contributed to flipped counties next.

Second, we can also consider the "flipped" counties as potential places where
disillusioned democrats had an impact. As Figure 1 showed, only 3 counties
flipped democrat -> republican (and 1 flipped republican -> democrat,
counteracting the other 3 somewhat, but we'll ignore that here). But let's
pessimistically attribute those losses entirely to democrats not voting for
Clinton. Combined, these 3 counties contributed 64,717 marginal votes to Trump
(only 17% of the 384,653 total votes Trump gained over Romney). Again, the total
impact is probably less, because these flipped counties also factor into the
some of the other phenomenons discussed in this section.

Finally, we look at how the changes in votes for just the democratic candidate
(**not** VD, just straight difference in votes Clinton and Obama received
correlates with changes in votes for independent candidates. If many democrats
opted to vote for independent candidates rather than Clinton in this election,
we should see a negative correlation between the two variables (indicating that
the increase in the independent vote is correlated the decrease in democratic
vote).

{{< figure src="/img/election_16/pa/dIV_by_dDems_16.png"
    link="/img/election_16/pa/dIV_by_dDems_16.png" width="100%"
    caption="Fig 6. Plot of dIV vs change in % votes for democratic candidate in 2012-2016. ">}}

However, Figure 6. doesn't confirm this theory. In fact, there's a slightly
positive correlation between the two values. This means that in counties where
Clinton loss the fewest votes (or even gained), percentage-wise, when compared
with Obama '12, there was a larger increase in the percentage of votes cast for
independent candidates.

But for argument's sake, let's assume that all of the additional votes (as a
percentage) the independent candidates received in 2016 (over 2012) came
exclusively from democrats who otherwise would have voted for Clinton. That
would explain only 142,027 of the marginal votes Trump garnered (37% of the
total extra marginal votes he won in 2016).

**Conclusion:** Although a popular refrain amongst Clinton-supporters in the
aftermath of the election, it appears that disenfranchised democrats didn't help
Trump as much as you might be led to believe. In the worst case, they explain
around 50% of the marginal votes he gained, but more likely they only explain in
the realm of 25% of them (assuming the independent votes came fairly uniformly from
Democrats and Republicans).

### Theory #3: Registered Democrats voted for Trump...

There's one last explanation for what happened - that registered democrats
decided to vote for Trump. We can't measure this directly from results and
registrations alone (we would need exit poll data to get a direct
estimate). However, when we compare the count of votes for Trump against the
number of registered republican voters, we find that in 14 of the counties Trump
won (remember, he won 56 counties) there were more votes cast for Trump than
registered republicans! And that is assuming that every registered republican in
those counties voted. Even so, that adds up to 56,575 votes for Trump that
couldn't have come from any registered republicans, representing 15% of the
extra marginal votes he gained over Romney.

**Conclusion** It's likely some Democrats (and Indepenents) voted for Trump, but
it's hard to say how many. It seems (I can't see any other explanation) that at
least 56,575 did, but it's likely more (hard to say how many more).

### Answer: It's Complicated

We started this section by asking "How did Trump win?" Here's a summary of our
answer:

1. Trump won a plurality of the vote in 56 (84%) of the counties. This was only
   2 more than Romney won in 2012.
2. However, Trump won 353,065 more marginal votes than Romney earned, enough
   that he beat Clinton by a mere 44,292 votes (less than 1% of all votes
   cast). In fact, that's 40,564 more marginal votes than Obama garnered in
   2012!
2. We can speculate about where these additional votes for Trump over Romney
   came from, but it's likely a combination of the following phenomenons:
 * Trump caused many Republicans to vote in the more rural counties. When
   comparing counties Trump won against counties Clinton won, we see a 3% higher
   turnout and 6% increase in RD. I'd guess this accounts for about 50% of the
   extra votes.
 * Trump flipped 3 counties, either because of above observed Republican turnout
   or because disillusioned Democrats stayed home. I'd guess this accounts for
   less than 17% of the extra votes.
 * Votes for independent candidates rose in almost every county. However, it
   seems unlikely that it this increase was due solely to Democrat voters. I'd
   guess this accounts for less than 37% of the extra votes (likely much less).
 * In 14 of the counties that Trump won, there were more votes cast for Trump
   than total registered republicans. This implies that some registered
   democrats (and independents) voted for Trump in at least those counties. This
   would account for at least 15% of the extra votes (likely more).

**Questions**

4. What would an analysis of precinct level results show?
5. How much of an impact did voter suppression laws have in PA?

<!---
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

**Careful - looks like turnout by party affiliation...**

***Methodology*** Figures 3 and 4 are histograms of the IV and turnout,
respectively, by county. Counties are first grouped by the party of the winning
candidate (blue for democrat, red for republican), and then binned.

***Discussion*** Looking at Figures 3 and 4, there's no obvious correlation
between the winning candidate in a county and either IV or turnout. Counties
where Trump won a plurality did have a higher median turnout, but the difference
is small.

However, this doesn't really answer the question either way. If anything, this
shows the answer is more subtle and complex. Perhaps "Bernie Bros" had an effect
in some of the more densely populated counties, but not in most of the largely
rural counties. Without looking at exit poll data and the demographics of the
population and their political preferences, it's unlikely we'll ever have an
answer to this hypothesis.
!--->

## Should we be surprised?

Okay, so we have examined the data for 2016 (and some from 2012) and have some
interesting results on what Trump did right. Nevertheless, many people (myself
included and, arguably, a majority of the US population) were surprised by the
election results. At this point I (and if you'd made it this far, I suspect you,
too) became curious: Is this surprising? Was this apparent sweep of counties
unprecedented? Afterall, we heard a lot about how this was the first year in a
long time that Pennsylvania "went red", so how did this year actually stack up
against previous elections?

To shed some light on these questions, we can take a more comprehensive look at
data from previous elections. This will help place this year's results in
context. I gathered county-level results for every presidential election in
Pennsylvania since 1996[^5], and registration data since 2000. In this section,
we examine several of the above reported metrics/variables from this wider
data set, to see if there are larger trends that have been at play in
Pennsylvania elections.

### Historical Election Results

We start, like we began the post, with a look at the voting results. This time,
from all elections since 1996.

{{< fig src="/img/election_16/pa/hist_results.png" >}} Fig 7. Three different
ways of visualizing the results of elections 1996-2016: Tukey boxplots of VD (left),
maps of VD (center) and breakdown of MV by candidate (right). {{< /fig >}}

While it's tempting to see Trump's victory as an anomaly (the "first time in
recent history Pennsylvania has gone red" perspective), Figure 7 tells a more
nuanced story. Indeed, the first thing you should notice is the gradual yet
steady shift in much of the state towards the right since the 1996
election. From this perspective, 2008 was itself an anomaly. It's the only
election when the VD actually shifted slightly back to the left and Obama won an
extraordinary number of additional MV while McCain lost around 100,000. In
comparison, Trump's win anomalous because he gained as much ground as Obama did
in '08, while the Democrats didn't lose *any* ground.

This trend is also suggested when examining flipped counties. With the exception
of Chester county in 2016, 2008 is the only year since 1996 when any county
flipped Democrat. That means that since 1996, a net of 17 counties have flipped
from voting Democrat (in 1996) to Republican (in 2016).

{{< fig src="/img/election_16/pa/hist_d_results.png" >}} Fig 8. Visualizing
changes in election results for years 1996-2016 with: Tukey boxplots of dVD (left),
map of dVD (right). {{< /fig >}}

***Methodology*** Figure 8 presents a map where every county is colored
proportional to its shift in VD (its dVD) from the previous year. To simplify
and better illustrate the general direction and magnitude of shifts, only 7
colors are possible. 3 shades of blue and red each, and white. The shades of red
progress form lightest to darkest, corresponding to 3 ranges of shifts: 3%-10%,
10%-20%, and 20%+. Similarly for blue. White indicates counties which had less
than a 3% shift in either direction.

I provide Figure 8 simply because it is the clearest way of seeing how
widespread the shift to the right has been in the state, even amongst counties
which continue voting Democrat. Also, I found it surprising how most of the
counties exhibit shifts of >3% every year, and most often in the same direction
for each election.

{{< fig src="/img/election_16/pa/results_IV_years.png" >}} Fig 9. Plotting IV
(left) and dIV (right) from the elections in 1996-2016. {{< /fig >}}

Finally, we wrap up this section by looking at how the independent vote has
changed over the years. Figure 9 shows two takeaways for me:

* Although IV was high in 2016, it was comparable to 2000 and is totally dwarfed
  by 1996.
* While IV increased in every county in 2016, this shift is insignificant in
  scale when compared with shift in VD (seen in Figure 8).

At this point, I am left with the following questions. These would be good
topics for further exploration.

**Questions**

* How does 2000 compare with 2016 (seems similar in structure - GOP made big
  gains while DEM stayed flat)?
* 2008 really does seem like an outlier, but in some ways 2012 is just as much
  an outlier - REPs made significant gains, and yet Obama still won...why?
* Is Fig 7 evidence that partisanship is growing? Normalize growth in total MV
  by population growth - seems like indicating increasing
  polarization/partisanship of state. That is, if, after accounting for
  population growth (isn't this VD?), we still see substantial increases in MV,
  that means that the same population is producing more MV (for one candidate or
  the other) per county, which suggests that the different counties are
  separating further from each other, becoming more polar.
* Why was IV so high in 1996? Had it been as high in previous elections? What
  happened between 1996 and 2000 that caused it to dropoff so dramatically?


### Historical Registration Data

Looking at results is one thing, but as we examined how Trump won earlier, we
saw that looking at registration data can also be helpful. To start, let's look
just at how the basic registration variables and how they've changed over the
years[^6].

{{% fig src="/img/election_16/pa/variables_dists_years.png" %}} Fig
10. Distributions by year of (from left to right): total registered voters (RV),
party registration differential (RD) and turnout. The second row shows how
each variable changed between elections. {{% /fig %}}

We can observe a few interesting things from Figure 10. Note that for most of
these, your guess is as good as mine as to what they signify...

* Total number of registered voters (RV) hasn't increased substantially since
  1996.
* dRD (differences in the Registration Differential between elections) has
  varied similarly to dVD (trending rightward, except for 2008). But the shifts
  are much less pronounced. Figure 11 explores this in more detail.
* Unsurprisingly, 2016 saw the highest turnout of the past 5 elections. However,
  I'm surprised to see that turnout dropped in 2008 - it seemed with how many
  extra MV Obama won, it would have been due at least in part to higher turnout.

{{% fig src="/img/election_16/pa/dvd_vs_drd_years.png" %}} Fig 11. Plotting
counties by how party registration (x axis) and vote (y axis) differentials
shifted between elections. Points are colored according to the winning party
(red=Republicans, blue=Democrats). {{% /fig %}}

Figure 11 dives into a point made above, that the distributions of dRD and dVD
look similar. As the figure shows, outside of 2012, they have a Pearson
correlation of 0.65 or better, which matches our intuition that they are
somewhat correlated. I find it interesting how 2008 and 20016 are similar
(mirror) to each other, while 2004 and 2008 are much more condensed. It's also
remarkable how few counties exhibit opposing shifts (shift right in RD but left
in VD, or vice versa), meaning dRD and dVD almost always have the same
sign. Also, it seems like `abs(dVD)>abs(dRD)`. I guess both are somewhat
intuitive, but that suggests that dRD is a decent first-order approximator for
VD: `VD = VD_old + dRD`. But I leave that for future exploration.

{{% fig src="/img/election_16/pa/registration_changes_years.png" %}}
Fig 12. Visualizing how the changes in registrations (dRV and dRD) between elections
distribute geographically (top) and along the two axis (bottom) {{% /fig %}}

I also wanted to explore how changes in total registration and party
differential jointly varied, which is what you see in Figure 12. Each point
corresponds to a county from the map above, and every county is colored
according to the 2D colormap on the right. The legend in the bottom right gives
a general way of interpreting the graphs based on which quadrant each county is
in. For example, counties in the upper left quadrant experience a net loss of
registrations but a shift towards a larger percentage of voters registered
Republican, meaning that more Democrats were unregistered than Republicans.

One thing that's pretty apparent is that dRV and dRD seem largely
uncorrelated. Otherwise, all I notice is that the southwestern counties are the
only ones to consistently see increases in relative percentage of Republicans
registered, while the southeastern counties have only gotten more
Democrat-leaning in registrations.

### Summary of all Data, by County

Finally, just because I can, I’ve plotted many of the variables presented above
for each county [here](/img/election_16/pa/counties_years.png). I find it pretty
interesting to look at each individual county to see how it’s changed over the
years. A very obvious curiosity is if one can train classifier on this data to
see if we can cluster the counties. I made some feeble attempts but didn’t find
any meaningful clusters.

## Closing Thoughts

First, thanks for making it this far! I started looking at the data used in this post
back in November, shortly after the election, and started writing this post in
early January. I've had a tremendous time trying to focus on producing a clean,
understandable set of results, and focus was difficult when the data is so
enticing. So I appreciate you taking the time to read the whole post.

I hope you read this all with an air of skepticism and welcome any and all
feedback. I'm not a statistician and didn't really pay attention to figuring out
if any of the results I was generating were statistically significant. Further,
as mentioned several times throughout the post, it's very difficult to discern
causation/explanations for the election results purely from the data I
examined. Nevertheless, I feel like I have developed a pretty good intuition for
what happened (at a macro level) in Pennsylvania, and what's been going on in
the state over the past several elections. But I have no context in which to
interpret this. I think an understanding of state demographics would be crucial
here. And of course, I'd love to see how this sort of analysis applies to other
states. Maybe someday.

<!---
***Questions***

1. Where there any counties where the VD was less than IV?
2. How does IV correlate with abs(VD)?
3. How much VD did each county contribute to Trump's victory?
4. How does 2016 compare to previous years?
5. How much is dVD accounted for by voters changing sides vs changes in
   turnout/registration?
6. Is 3% VD shift statistically significant?
7. Can we cluster the counties based on the data (VD, dVD, VD-median,
   dVD-median, turnout, RD, etc.)?

{{< figure src="/img/election_16/pa/results_map_years.png"
    link="/img/election_16/pa/results_map_years.png" width="100%"
    caption="Fig 7. Map of election results for years 1996-2016." >}}

{{< figure src="/img/election_16/pa/shifts_map_years.png"
    link="/img/election_16/pa/shifts_map_years.png" width="100%"
    caption="Fig 8. Map of dVD by county for elections 2000-2016." >}}

{{< figure src="/img/election_16/pa/results_barh_years.png"
    link="/img/election_16/pa/results_barh_years.png" width="100%"
    caption="Fig 9. Marginal votes by winning party for years 1996-2016." >}}

{{< figure src="/img/election_16/pa/results_VD_by_RV_12-16.png"
    link="/img/election_16/pa/results_VD_by_RV_12-16.png" width="100%"
    caption="Fig X. Shifts in VD by RV in each county. Each arrow represents how a county's VD and RV changed between 2012 and 2016." >}}
!--->

[^4]: I'll often omit the sign for MV, opting to instead report the candidate
    who gained those marginal votes instead.

[^1]: A more common practice is to use voting age population (VAP) or voting
    eligible population (VEP) as the denominator when computing turnout. VAP is
    readily available from the US Census, and so is a popular choice, but has
    many flaws, especially when comparing across states as different states
    place different restrictions on who is actually eligible to vote within the
    VAP. VEP is a fairly recent idea, and is meant to account for all the
    differences in states and thus be more readily comparable between
    them. However, it relies on data that is not always easy to find (or even
    public at all) at the county level, and is thus nearly impossible for me to
    use in this analysis. Michael McDonald has good set of posts
    on [VAP](http://www.electproject.org/home/voter-turnout/faq/vap)
    vs [VEP](http://www.electproject.org/home/voter-turnout/faq/sold) and
    their
    [differences](http://www.electproject.org/home/voter-turnout/faq/vap-v-vap)
    and
    why
    [using registration is not ideal](http://www.electproject.org/home/voter-turnout/faq/reg). In
    all honesty, the biggest flaw with using RV instead of VAP in this analysis
    is that I can't separate changes in registration due to population
    growth/change vs political views shifting.

[^2]: Measuring difference/change is finicky when using normalized variables, as
    I've realized. In this piece, I've tried to consistently compute and present
    differences as differences between measurements. But when these measurements
    are normalized (e.g. IV=[votes for independent candidates]/[total votes
    cast]) (and usually interpreted as percentages), there
    are
    [many seemingly valid](https://en.wikipedia.org/wiki/Relative_change_and_difference) alternatives
    for computing differences. I'm not really convinced that any of them is
    better than the other, but I also don't profess to being a
    statistician. Just let the record state that this is one ignorance I am
    aware of.

[^3]: It's interesting to note that this is*not* the opposite of "How did
    Clinton lose?", although the two are obviously closely related. It turns out
    that Trump and the GOP did many things (or had many things go right,
    depending on your perspective) to win, and presents a more interesting
    analysis than the "hindsight is 20/20" sort of analysis of why Clinton
    lost. Also, as you'll notice from the data, Clinton garnered nearly the same
    votes and turnout as Obama did in '12, so we'd really just be looking at all
    the places where Clinton failed to predict/counteract Trump's successes.

[^5]: Why 1996? Somewhat arbitrary, although it gives us context from the last
    two party shifts in the presidency (Bush in '00 and Obama in '08). Also, the
    Pennsylvania state website only hosts results since 2000, so I had to get
    creative in my google searches. Finally, I hope one day to run this analysis
    for other states, and the farther I go back, the more data I'd need to
    format and parse to analyze other states (where I can even get it at
    all). Call me lazy.

[^6]: As mentioned before, 1996 is missing when we consider registration data
    because I couldn't find it at the county level.

[^7]: It's not, due to the way the electoral college works in America. I suspect
    the general population's understanding of the electoral college spiked
    during and after the election this year.
