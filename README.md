# CARDS API DEMO

### Description
Basic REST API written in GO-lang with couple of endpoints which give back a whole deck of playing cards or a specific number of cards drawn from a deck. A deck could be issued sorted or shuffled. Once issued deck could be found by its ID and shuffled as well. Cards could be drawn from each issued deck (from 1 card to the number of cards left in a deck) or from a new (shuffled or not) deck.

#
### Endpoints
/api/deck/new/ </br>
/api/deck/new/shuffle </br>
/api/deck/new/draw/count/{count} </br>
/api/deck/{id}/get </br>
/api/deck/{id}/shuffle </br>
/api/deck/{id}/draw/count/{count} </br>

#
