## Description

Building a `personal` AI travel assistant. The aim of this AI assistant is to have friendly conversations with people while helping them build and manage their travel plans. Pricing Model: TBD

## Functions

The following are some of the functions of this AI model:

- Recommend hotels based on the user's budget with links for booking etc
- Recommend tourist attraction based on the user's budget and interest with links for booking etc
- Tell the user about the weather and the best period to visit these locations
- Recommend flights based on the user's budget with links for booking etc
- Recommend restuarants and help plan dates with relevant links for booking etc
- Recommend places to go shopping with discounts if available
- Inform the user about natural disasters and extreme weather conditions in places they visit
- Plan itenaries andn input them in calendars
- Handle payments with user confirmation i.e Human in the Loop (LATER)

## Frontend Requirements

The following are some of the requirements of the frontend app:

- Server rendered JS app, nothing complicated. So we use Remix
- App should be a PWA app so users can easily install it
- App should be fast. AI is naturally slow so we have to make up for this
- Interface should be mega simple
- Use user & agent name when constructing chat interface
- Buttons for having conversations on whatsapp, telegram (LATER)
- Add support for multiple languages

## Backend Requirements

The following are some of the requirements of the backend app:

- Should be super fast so we will be writting this in golang
- Use concurrency whereever possible
- Should be able to stream response to frontend
- Interrupt long running request and tell the user that we will get back to them
- Invite others into the chat (LATER)
- Voice call (LATER)

## AI

The following are some of the requirements of the AI:

- This AI needs a persona so we will build one
- We will have a manager Agent which interacts with the user. Gets information and communication with other agents to execute task
- Each agent will be built for a particular purpose. E.g A single agent will handle everything bookings and another can handle everything flights related etc
- If an agent needs more time to complete its' task, it needs to communicate this to the manager Agent which will inturn communicate this to the user
- AI will never compromise user so it shouldn't store user important information
- If AI cannot do something then it clearly tells the user e.g payments etc
- AI should make a short introduction asking the user what name it should be called and what name it should address the user as
- Messages should not be too long or overwhelming
- Support for voice notes for messages which are a bit longer
- AI should remember user through multiple chat context

## TODO

- Add icons to distinguish places, attractions, hotels and flights

- Formatting the response properly and make the clicks work [maps, look and fee etc]
- Integrate with actual data
- Work on pricing and other miscellaneous
