# repository interface

the repositories are the abstractions for resources

mysql, rpc, kafka, redis are all resources

but cache is not, cache is a cross-cutting concern

https://stackoverflow.com/questions/7009306/is-caching-a-repository-domain-or-application-concern

if you use redis as cache, it'll also be a cross-cutting concern.
