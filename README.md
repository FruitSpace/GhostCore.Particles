# GhostCore Modules



### VM Pooling

1) Core starts N javascript VMs in a global pool to maintain availability
2) Each VM gets populated with GCRuntime and other modules
3) Each GDPS gets assigned with a smaller pool of VMs from the global pool every request
4) Modules are loaded from cache separately in different runtimes
5) Each runtime gets a timeout of 30s (adjustable)
6) After there are no more js and messaging events, VMs get "cleaned" and returned to the global pool
7) If timeout is reached, VMs are killed

### Messaging and events
See [Runtime Readme](runtime/README.md)