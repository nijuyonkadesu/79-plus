## spec: (iteration 01)

~~v1/user/(id)/chat~~
~~v1/channel/(id)/listen

No, do not use the above url path. Coz, they're rest style resource. Instead, use a single ws path, where all routing happens.

```lua
POST    v1/messages
GET     v1/ws
```

- users will be able to post message on either medium when they have proper permissions
- 1 client = 1 ws connection

TODO: what is room concept in gorilla websockets?
TODO: reliable notifications & retries
TODO: tag users & notifications (extension of channels)
TODO: group chats (extension of channels)
TODO: refer api structure of telegram
TODO: userbots (preferrably, I'll reuse tgbotserver & tgbotapi as is to save some implementation time?) - but first, do a crude version of bot api, discard it, use tgbot api & server
TODO: fix celebrity problem

### thinking

TODO: re-iterate on the hub design (as a concurrency boundary), and as a design choice for the chat application. (I just need future implication and the expected changes at the later stage)
```sh
# Structurally:
[ws] -> room
[room] -> hub # like connection pool, this is going to be 'room pool' ? to avoid creating too many rooms?? (indirectly exhausting unix sockets)

# Functionally: 
ws -> hub -> room [create lazily for each dm / channel / group when their users come online??] at join stage. # we'll figure scale later, coz idk how to do 
```

a room ensures:
room.clients = only-online-and-allowed-clients

joining a client to a room does: 
- load membership
- check permission
- check mute
- check presence
room moves out heavy logic out to join/leave stage, not send time.

TODO: slow clients will kill throughput

The data processing part: (stateless)
1. auth
2. read the response
3. process the payload with the right handler based on message_type
4. write back the response

Someone needs to maintain the state... (the room), bu

## flow: (iteration 01)

1. upgrade to websockets
2. get (n-x) old unread messages
3. receive new messages (type: message)
4. post new messages
5. (later) fanout / multiplex writes to both queues and db (scylla db & es?) - should be optimised for large volume of instant writes
6. (later) partition and shard database

- 3 & 4 functions by listenting and pushing to a queue (redis or other mq)

GPT:
1. HTTP GET /v1/ws → Upgrade
2. Client sends "auth" message
3. Server validates, attaches userID
4. Client sends "subscribe" to rooms
5. Server sends last N unread messages
6. Bi-directional real-time messaging

TODO: websocket reconnection and consistent hashing
TODO: encryption & searchable text
TODO: media processing
TODO: does protobuf be of any help here?

## Design goals:

- should not be a microservice - A monolith, but the feature-flags should turn off unnecessary features - so it can be deployed as a microservice with muliple configurations if needed.
- at later stage (ITR 3 or 4) - think about how to make the system more friendly for backward and forward compatibility.
