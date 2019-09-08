# Decorator Pattern 

Decorators are used to insert additional functionalities around functions/methods without modifying it. For example, a
common usecase could be to log the time taken for a function to execute such as a HTTP client call or something equivalent.

We create a wrapper around such client function and mark the time before and after the function executes and then calculate the 
difference. This way any function can be measured without having to modify its original logic. 