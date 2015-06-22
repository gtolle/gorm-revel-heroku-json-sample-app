Here's a sample (but realistic) Go webapp that provides a JSON API using Revel and Gorm to connect with Heroku Postgres through their standard DATABASE_URL mechanism.

I just started learning Go, and found this app helpful to create.

I come from a Rails background, and while it takes a lot of work to create something as full-featured, powerful, and flexible as ActiveRecord, I'm pretty darn impressed with Gorm so far. Preloading is one of my favorite ActiveRecord features, and I'm so glad that Gorm has it.

I'm also pretty impressed with Revel, and Go's JSON marshalling.

Hope you find this helpful for starting your own apps too.

For deploying to Heroku, I'm using the Revel Buildpack. To use it, you just have to create your Heroku app like this:

```
heroku create -b https://github.com/revel/heroku-buildpack-go-revel.git
```

Then you can just do the usual `git remote` stuff and `git push heroku master`.
