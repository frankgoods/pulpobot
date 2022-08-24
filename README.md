# pulpobot

Your friend is going through hard times? Your wife is far away and you want to support her with all your heart? Or maybe it is you, who needs to hear something nice?

**Pulpobot** is here to help!

What it does is simple, yet so powerful! You provide the bot with some sample text set, and it's going to send randomly created phrases to your subscriber!
The bot is very special so it is only one subscriber that can be served. You create something special for someone special, right?

## Usage
Bot is very simple, very human.

All you need is right here:
1) Create a telegram bot, I assume you know how to do it (if not, just look here https://core.telegram.org/bots/)
2) Place your bot token to `priv/token` file
3) Change the secret password in the `priv/password` file. By default it's just the word `pass`
4) Run the bot!
5) Share the secret password with your friend or husband or whoever is going to enjoy your bot's company (but we know, it is you they going to enjoy!)
6) Tell your dear friend to send the following command `/start <your secret password you just created in the item 3>`
7) That's all! Bot is going to recognize its target subscriber and start a conversation with her/him.

## Tailoring bot's text samples
You might wanna change bot's default text samples and create something creative and intelligent.
Check out the `priv/talk` file.
The structure of the file is very simple but I'll comment on it a bit anyways:
```
{
    "Begin":
        ["Friend, ","Bro, ","Hey!],
    "Middle":
        ["all good ","you have a problem ","just do it "],
    "End":
        ["and relax","but be yourself","but be a man","this is bullshit!","and forget about everything"],
    "Whole":
        ["Your life, your choice","I feel bad for you!",
        "Nonsense...And what's more, it doesn't rhyme. All decent predictions rhyme"
        ],
    "ReplyText":
        ["Ok", "Accepted", "Got you!", "Thank you"]
}
```
Here is how it works. Bot will pick randomly something from `Begin`, `Middle` and `End`, combine these into a sentence and send it.
Sometimes it will just send something from the `Whole` array.
And if your subscriber is kind enough to send something back to the bot, our bot will reply with something from `ReplyText`

That's all you need to know!

## One more thing
Bot sends a message to a subscriber once in a while. You can control how often this happens. You've got two command line arguments you can supply to your bot:
1) Minimum delay between two messages. Use, for example, `-min 15` to send a message every 15 minutes
2) Additional random delay your bot gonna wait on top of minimum delay. Use something like `-rand 20` to make the bot wait for another 0 to 19 minutes.

The default values are `-min 35` and `-rand 30`.
