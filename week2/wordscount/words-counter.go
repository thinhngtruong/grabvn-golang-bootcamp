package wordscount

/*
Requirements:
- Input: file paths, folder paths, urls (must have http or https prefix) or both 3 kinds of them. 
- Input values are separated by commas.
- Only process 10 resources at a time (use worker pool).
- Time out for each resource is 10s.
- Wait for all resources processed, then show result.
*/