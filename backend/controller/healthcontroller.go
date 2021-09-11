package controller

import (
    "context"
    "fmt"
    "NewsAPISPE/backend/config"
    "NewsAPISPE/backend/models"
    "log"
    //"errors"
    "database/sql"
    //"strconv"
)
 
const   tablehealth = "health"
// ambil semua data dari database
func GetHealth(ctx context.Context) ([]models.Article, error) {
 
    var articles []models.Article
 
    db, err := config.Mysql()
 
    if err != nil {
        log.Fatal("Gagal terhubung ke MySQL", err)
    }
 
    queryText := fmt.Sprintf("SELECT * FROM %v", tablehealth)
 
    rowQuery, err := db.QueryContext(ctx, queryText)
 
    if err != nil {
        log.Fatal(err)
    }
 
    for rowQuery.Next() {
        var article models.Article
 
        if err = rowQuery.Scan(
            &article.Id,
            &article.SourceId,
            &article.SourceName,
            &article.Author,
            &article.Title,
            &article.Description,
            &article.Url,
            &article.UrlToImage,
            &article.PublishedAt,
            &article.Content);
			err != nil && sql.ErrNoRows != nil{
            return nil, err
        } 
        articles = append(articles, article)
    }
 
    return articles, nil
}

// get selected
func GetSelectedHealth(ctx context.Context, art models.Article) ([]models.Article, error) {
 
    var articles []models.Article
 
    db, err := config.Mysql()
 
    if err != nil {
        log.Fatal("Gagal terhubung ke MySQL", err)
    }
 
    queryText := fmt.Sprintf("SELECT * FROM %v WHERE id = '%d'", tablehealth, art.Id)
 
    rowQuery, err := db.QueryContext(ctx, queryText)
 
    if err != nil {
        log.Fatal(err)
    }
 
    for rowQuery.Next() {
        var article models.Article
 
        if err = rowQuery.Scan(
			&article.Id,
            &article.SourceId,
            &article.SourceName,
            &article.Title,
            &article.Author,
            &article.Description,
            &article.Url,
            &article.UrlToImage,
            &article.PublishedAt,
            &article.Content);
			err != nil && sql.ErrNoRows != nil{
            return nil, err
        } 
        articles = append(articles, article)
    }
 
    return articles, nil
}