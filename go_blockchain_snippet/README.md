> Simple golang implemented blockchain to keep bookstore transaction logs

                cd go_blockchain_snippet
                go mod init github.com/Jasmeet-1998/go-orientations/go_blockchain_snippet

                go mod tidy

1. Book

- ID
- Title
- Author
- PublishDate
- ISBN

2. BookBoughtTicket [Information stored in the blockchain] i.e the DATA section of the block

- User
- DateOfBuy
- IsGenesis (boolean)

> NOTE- the first block in blockchain is genesis block and it does not have PrevHash field.

> Block

- reversed linked list

        |---------------------------------------------|
        | PrevHash| Position| DATA | TimeStamp| Hash| |
        |---------------------------------------------|

- PrevHash points to the hash of the block just before the current block
- Hash is the sha256 based hash for the current blockchain that act as refference point for the subsequent block to have the PrevHash reff to this Hash field.
- position is position of the block in the blockchain and Timestamp is when the block was created or updated.

> 📝 Development

1.  struct Book
2.  struct Block
3.  struct BookTicket
4.  struct Blockchain

        # functional flow
        Book -> BookTicket ->Block -> Blockchain

> Blueprint/Walkthrough

        # to create new blockchain
        New Blockchain --> Write Block --> Add Block --> Create Block-> check Genesis or not -> validate block-> generate hash-> validate hash

> run/test locally

                    cd go_blockchain_snippet
                    go run main.go

                    # add new book
                    # http://localhost:5000/new
                    {
                        "title":"Book1",
                        "author":"Anonymous1",
                        "isbn":"345678",
                        "publish_date":"1909-05-07"
                    }


                    # copy the newly creted book id from response
                    # and create new block -> http://localhost:5000 (POST)


                    # get current blockchain
                    http://localhost:5000 (GET)
