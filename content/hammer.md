## How to install hammer editor on linux

Although valve is known for good linux support, there are things, they didn't port yet like Alien Swarm or their SDK. A while back, I wanted to get back to making maps, but I don't dual boot windows anymore, so I had to come up with a way to install hammer on linux.
Here I will show you, how I did it.
Hammer actually comes with valve games, but only on windows. If we want to install it on linux there are two ways:
1. download and install windows version of steam thru wine and then install a valve game
2. download windows binaries for a valve game

In this tutorial, I will show the second one.



### prequisities
For this to work you will need: steam, some valve game (recommend tf2 or cs:go), wine and wine32.

### step 1: download hammer
To download hammer, we will need to download windows binaried for a valve game (i will show it on tf2). We will use steam console to achieve this. To open steam console write `steam://console` into your browser.
Then if you head to your steam client you will see console button on the right of your profile name. When you have the console open, you can use the `download_depot <game_id> <depot_id>` command to download depots.
You can look up game id on [steamdb](steamdb.info). For tf2 it is `440` and we want to download the `TF2 Windows client` depot which has number `232253`. That means the following command will look like this: `download_depot 440 232253`.
It will take a while to download

### step 2: copying
When the download is finnished, the downloaded data will be in `~/.steam/steam/steamapps/content/app_<game_id>/depot_<depot_id>/`.
To use this copy content of the depot folder to a game folder, commonly placed in `~/.local/share/Steam/steamapps/common/`.

### step 3: running
After copying the files, we can finally launch hammer. It is located in `~/.local/share/Steam/steamapps/common/<game_name>/bin/hammer.exe`. We will use wine to launch it.
It should start normally. If not the errors are usually easy to solve by just googling.

### step 4: make maps



Thanks for reading this guide. Hope it works :]
