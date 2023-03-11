Low Code NewsLetter
====

### Create new resource named NewsLetter
Create new yml file (news-letter.yml) for describing newsletter resource
```yml
name: NewsLetter
sourceConfig:
    dataSource: default
    entity: news-letter
properties:
    - name: email_address
      type: STRING
    - name: category
      type: STRING
```
Now let's create resource with given 
```shell
dhctl apply -f news-letter.yml
````

### Generating Golang codes for created resources
```shell
dhctl generate NewsLetter --path=model
```
**You will get following structure created**
```
type NewsLetter struct {
	Id          string
	Name        string
	Description string
	Version     uint64
	CreatedBy   string
	UpdatedBy   *string
	CreatedOn   time.Time
	UpdatedOn   *time.Time
}

```
### Implementation of send mail when newsletter is saved

```
func main() {

    client := getDhClient()
    
    newsLetterExt := client.newExtension(&model.NewsLetter{})
    
    newsLetterExt.AfterSave(func(newsLetter *model.NewsLetter) {
        // implement send mail codes
    })
    
    
    newsLetterExt.Run()
}

```

### deployment of code

```shell
dhctl deploy newsLetterExt
```