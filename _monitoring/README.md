# Monitoring

How do we actually know what is being put onto the ledger, and is the idemix configuration actually hiding the identity.
One way is to do block eventing monitoring - and look at the blocks that are addedd to the ledger and decode the contents.
(this doesn't violate security as you need to have an admin identity to do this)


## Application

This uses the fabric-gateway's nodejs client to listen for and decode block events. It's been hard-coded to use the identities
and connections created in the test app (in the directory above this)

## Build

Launch a new shell to run the blockevent listener

```
```

## Run

