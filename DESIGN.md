## Councillors and Locals.

Both councillors and locals are users of the system.
An admin can create councillors within the system, this creates a user that cannot sign in but will
show up for locals.

When a councillor signs up, they should use their public email. As we should already have this registered,
we can then send a mail to that email address looking for verification. When the email is verified, then the
verified councillor can set a password and will be able to login as normal.


Current domains:

- locals 
    - encompasses local residents and local councillors
- communications
    - handles the different types of communications that can happen between users and councillors 
- projects
    - projects represent a group effort to try and effect change within a local community. 



## Refactors to complete

Remove admin countroller and integrate create councillor into the councillor controller with an admin:scope access

Remove duplicated counties and areas in client 

sort out the councillors area of the client. It is disperate and needs more cohesion 


## communications

A local constituent is the only one who can initialise communication and they can only communicate with a councillor that is within their area.
Once communication has been initialised, then the councillor is free to respond.
Only the initialisor of the communication can decide that the communication is complete.