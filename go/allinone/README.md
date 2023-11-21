# APIs
- AddSinglePersonAndMatch
  - Add a new user to the matching system and find any
possible matches for the new user.

- RemoveSinglePerson
  -  Remove a user from the matching system so that the user
cannot be matched anymore.

- QuerySinglePeople
  - Find the most N possible matched single people, where N is a
request parameter.

# Here is the matching rule:
- A single person has four input parameters: name, height, gender, and number of
wanted dates.
- Boys can only match girls who have lower height. Conversely, girls match boys who
are taller.
- Once the girl and boy match, they both use up one date. When their number of dates
becomes zero, they should be removed from the matching system.

# Other Requirements
- Unit Tests
- Docker Image
- Structured Project Layout
- API Documentation
- System Design Documentation That Also Explains the Time Complexity of Your API
- TBD Tasks