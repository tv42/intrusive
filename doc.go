/*

Package intrusive contains intrusive container data types.

An intrusive data structure is one where the data items contain the
metadata needed, instead of the metadata pointing to the contained
data.

For example, a naive / pointing single linked list:

    -> Item.next -> Item.next -> nil
           .obj         .obj
             |            |
             v            v
            data         data

Intrusive linked list:

    -> data
           .Item.next -> data
                             .Item.next -> nil

The goal of this is to decrease the number of allocations; instead of
`Item` and `data` being separate chunks of memory, they are allocated
as one.

For more on this, see

    - http://www.codefarms.com/publications/intrusiv/intr.htm
    - http://kernelnewbies.org/FAQ/LinkedLists

Status

This is an EXPERIMENTAL package. Use at your own risk.

*/
package intrusive
