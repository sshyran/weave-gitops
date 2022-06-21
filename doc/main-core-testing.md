# Main core testing

## What is this?

This is emphatically not an ADR or an RFC. Both of those documents a
decision - this is not a decision, and merging it will not cause
stories to match reality to the text here to be created. This
documents a few steps that Robin would like to explore further, to
reach a place where he thinks we'll be better off. If Robin mumbles
about "long term testing" or "fixing testing", he's thinking about
something like this.

## Context

To make gitops a robust, quality application, it is important that we
can test it in a variety of ways, and strike a balance between ease of
writing and maintaining tests, time to run tests, and variety of
situations tested. That exact balance will always be evolving - this
document tries to stake out an architectural path to set
us up for success.

A major source of pain right now is that there's no distinction
between method and content, which makes it slow and frustrating to
build tests.

"Method" refers to the fact that our database is etcd which we
can only access through k8s APIs - they are good APIs, but they don't
provide all the generic, reusable web functionality that a mature ORM
for an SQL database will give you out of the box. So we have a lot of
code to implement the method to access data.

"Content" refers to the data itself, e.g. the objects and their
schemas. The k8s schemas that define the documents that we read are
quite stable, but they do expand, new versions are released, and the
data gitops want to display might not be trivial to get hold of.

We test "content" and "method" at the same time by spinning up a k8s
environment to talk to. This makes the tests slow to make the method
realistic, but it also limits the kind of content we can test by what
the k8s environment will return to us.

## Proposal

Let's move the k8s object interpretation to the frontend.

The core API will be shrunk down significantly - `getObject` (which
exist) for detail pages, `listObjects` (which does not) for table
views, and possibly a more restricted endpoint or two for non-flux
objects.

This means a language barrier between method (go code) and content
(javascript code), which we should make use of to separate the tests.

### Method

The backend can focus on just the method.

For backend development, it reduces the complexity in the code - all
test code related to making sure the backend interprets k8s data can
just be deleted (or rather, ported to javascript). Instead, the
backend will have a very small set of very generic endpoints. This
means the same number of backend programmers and backend tests will be
sufficient for a deeper, more robust test suite, as well as shipping more
features, faster.

It means that the backend can focus on building a robust database
layer. It means lots of UI changes can swoop by without backend
involvement, so the backend can focus on robustness, performance, and
new functionality, instead of constantly being interrupted with small
UI requests.

### Content

The frontend can handle, test, and display the content much more
easily on its own.

It means that for any problem we see, we can just grab a json or yaml
document that's problematic, stick it in a directory, and use it
directly in both unit tests and snapshot tests - and they can be both
old schemas and new schemas, without requiring a cluster running that
version. In other word, our tests for content can be much deeper, much
better tested, while running much faster.

It does mean taking into account some of the complexity that the
backend has previously dealt with. It also means doing this without
being able to re-use helper methods from go libraries.

But the flip side of that coin is that the user-facing developer has
the strongest connection to what data we need, when and where we need
it, and how to display it. Being able to ask the backend for just "all
the data" and dealing with it locally means that that the frontend
developer won't have to go through the cycle of "describe the problem
to a backend programmer, wait for it to be implemented, look for an
API update, use that API, see if it's correct" - not having to do that
is a big boost in productivity, even when the frontend and the backend
developer is the same person.

## Out of scope: acceptance testing

There's a [separate document](./acceptance-testing.md) for acceptance testing.
