# Core acceptance testing

## What is this?

This is emphatically not an ADR or an RFC. Both of those documents a
decision - this is not a decision, and merging it will not cause
stories to match reality to the text here to be created. This
documents a few steps that Robin would like to explore further, to
reach a place where he thinks we'll be better off. If Robin mumbles
about "long term testing" or "fixing testing", he's thinking about
something like this.

## Context

In [Main core testing](./main-core-testing.md), a mechanism to ensure
that gitops works well in a single environment was described. However,
we support many environments, so we need to test across them all.

Currently, there is no testing - automated or otherwise - for any
environment other than the main development environment, and our one
staging environment.

However, after release we need to support the application across a
variety of environments - across flux versions, k8s versions, and
cloud providers. There is a section in the gitops manual that
describes which versions of k8s and flux are supported, but it's
manually maintained by Robin using a thoroughly incomplete
makefile-calls-python-calls-makefile hack that nobody deserves to
suffer through.

## Proposal

We'll create an acceptance test suite for all tested
environments. However, it will be as simple and as feature-free as
possible - all acceptance tests are "annoying", so we'll avoid
building it out as much as we can.

 * All supported content schemas are quickly and frequently tested by
   the UI. So either UI works, or the UI test suite needs exending.
 * All supported backend features are quickly and frequently tested by
   the backend. Either the backend works for the k8s environment it's
   running, or the backend test suite needs extending.
 * We now know the UI works _if the backend works_. We know backend
   works _in 1 environment_. Now we need to test "do all environments
   work".

The general approach will be

 * Request URL from known-good environment.
 * Save response.
 * Request same URL from another environment.
 * Check that the response is the same.
 * Go to start.

We shouldn't ever be doing this for more than a dozen or so URLs.

To decide on which environments to use, we'll just do simple property
testing: we'll create a table of properties and supported values, such as:

 * Supported k8s versions
 * Supported flux versions
 * Supported optional k8s features
 * Supported cloud environments

For each row, grab a random valid value, spin up the environment, and
run the tests. Keep running these tests one at a time until we're
bored. Do this regularly, so the probability of bugs sneaking through
is kept low. Do this a little bit extra before an RC is tagged stable.

This test suite should be too simple for its failures to be directly
actionable: it should not tell you "there's a problem with
pagination", it should just tell you "when I ran an older k8s, and
queried `?page=2`, it's different - this is how you spin up your own
environment to test for yourself" and then you can spin that up
yourself and have a look.

Note that it's implied here that this test suite won't be a tool we
like using - it's supposed to be annoying and regularly break for
invalid reasons and provide rubbish feedback. This is how we remember
not to extend this suite, but rather build out the quick, easy,
sensible tests, because those are the ones that give us fast, good
feedback.

## What not to do: click testing

This suite does our acceptance testing, not click testing such as
selenium.

If the API response is the same, then the frontend will be fed the
same input. If the frontend is fed the same input, it will render the
same page. So if this suite passes, then the same test case as a
selenium test would also pass. Any desire to add a separate test suite
for selenium should be channelled towards improving our frontend test
suite.

## Out of scope: benchmarking

We should have a known, somewhat static environment, and regularly
benchmark performance to see that we don't have performance
regressions, and that our performance fits within acceptable
boundaries.

That is also part of acceptance testing, this is something we should
have, however that is out of scope for this test suite.
