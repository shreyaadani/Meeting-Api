
# Meeting-Api by shreyaa dani
a basic version of meeting scheduling API buid using Go

It executes folowing services:
 <li>Schedule a meeting:</li>
   <ul>   ○	it is POST request
      ○	Uses JSON request body
      ○	works on http://localhost:9090/meetings
      ○	returns the meeting in JSON format </ul>
<li>Get a meeting using id</li>
<ul>
○	it is a GET request
○	Id is taken as the url parameter
○	works on http://localhost:9090/meeting/id (where id= meeting id)
○	returns the meeting in JSON format </ul>
<li>Lists all meetings within a time frame</li>
<ul>
○	its a GET request
○	it works on http://localhost:9090/meeting?StartT=str1&EndT=str2 (where the start and end time is inputed in str1 and str2 respectively)
○	returns meetings in JSON format that are within the time range</ul>
<li>Lists all meetings</li>
<ul>
 ○	its a GET request
○	it works on http://localhost:9090/all 
○	returns all meetings in JSON format 
</ul>


<hr>

![ALL THE OUTPUT IMAGES](https://github.com/shreyaadani/Meeting-Api/issues/1#issue-724385005)
