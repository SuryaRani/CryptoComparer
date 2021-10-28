# Questions
1. Are there any sub-optimal choices( or short cuts taken due to limited time ) in your implementation?

2. Is any part of it over-designed? ( It is fine to over-design to showcase your skills as long as you are clear about it)

3. If you have to scale your solution to 100 users/second traffic what changes would you make, if any?

4. What are some other enhancements you would have made, if you had more time to do this implementation


# Answers
1. Yes,100% I was able to get a little bit of what I had in mind into what is in front of you right now, but given more time I would have added:
    * Making the website in React to be able to have the prices constantly updating and displaying instead of having to refresh to get the new prices
    * Probably play with the colors and html design a little more to make it prettier and just look nicer in general
    * Pictures that update with the prices for each exchange and for whether its buying or selling.

2. Honestly I dont think so I believe that I tried to be concise in my decisions to be able to show what I am capable of, and also make the app look half decent. Given more time I would have definitely allowed myself to spend alot more time on the front end design of the app. 

3. 
    * My applicaton is very lean, but one thing that could happen with that many requests is that I would get rate limited by my API calls. What I would have to do is most likely pay for an upgraded version of the API's I use to be able to increase the amount of requests in a certain time frame I could make. 
    * Since I am hosting my app on Heroku I would also have to make sure that it could handle the traffic I am recieving. I would most likely also need to pay a premium fee to be able to service that many users and requests per second. 

The rest of my program is extremely lean and can be scaled upwards easily as it is only a static html page and Go making a few API calls in the back end. 

4. As stated before probably the biggest thing I would have wanted to add was dynamic price change on the website. Especially with cryptocurrency and how volatile it can be, one must be able to show crypto prices live at any given time, which is one big problem in my implementation. I would have loved to have written this in React, but I have never written in React before and although under normal circumstances I definitely would have learned it, with this small time frame and midterms coming up I decided that it would be better for me to make it a little simpler. 

# Access Live Version
* If you want to see the live version of this program it is being hosted on Heroku 
    * https://crypto-comparer.herokuapp.com