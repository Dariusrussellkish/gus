namespace go message

struct Message {
    1: i32 id,
    2: string to,
    3: string frm,
    4: string content,
}

struct MessageAck {
    1: i32 id,
    2: string frm,
}

service Messenger {

    MessageAck sendMessage(1: Message m)

}