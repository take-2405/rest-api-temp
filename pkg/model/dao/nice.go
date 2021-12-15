package dao

import (
	"backend-record/pkg/model/dto"
	"database/sql"
	"github.com/google/uuid"
	"log"
)

const(
	getNices                  = "SELECT nice FROM articles_contents WHERE article_id=?;"
	updateNices               = "UPDATE articles_contents SET nice = ?  WHERE article_id=?;"
	insertArticleNiceStatus   = "INSERT INTO articles_nice_status VALUES (?,?,?);"
	deleteNiceStatus          = "DELETE FROM articles_nice_status WHERE user_id=? AND article_id=?;"
)

// /update/add/like
type upDateAddLike struct{
	ArticleID string `json:"articleID"`
}

func MakeUpdateAddLikeClient(request dto.NiceRequest)upDateAddLike{
	return upDateAddLike{ArticleID: request.ArticleID}
}

func (info *upDateAddLike)Request(userID string) (*dto.Nice, error) {
	var nice dto.Nice
	var	err  error
	//nice取得
	row := Conn.QueryRow(getNices, info.ArticleID)
	if err = row.Scan(&nice.Nice); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		log.Println(err)
		return nil, err
	}
	status, err :=CheckNiceStatus(userID,info.ArticleID)
	//場合分け
	if status == false { //追加
		uuid, err := uuid.NewRandom()
		if err != nil {
			log.Println("NiceID is error")
		}
		niceID := uuid.String()
		_, err = Conn.Query(insertArticleNiceStatus, niceID, info.ArticleID, userID)
		if err != nil {
			return nil, err
		}
		nice.Nice = nice.Nice + 1
	} else { //削除
		_, err = Conn.Query(deleteNiceStatus, userID, info.ArticleID)
		if err != nil {
			return nil, err
		}
		nice.Nice = nice.Nice - 1
	}
	//nice更新
	_, err = Conn.Query(updateNices, nice.Nice, info.ArticleID)
	if err != nil {
		return nil, err
	}
	return &nice, nil
}

func CheckNiceStatus(userID,ArticleID string) (bool, error) {
	var niceID string
	status := false
	row := Conn.QueryRow(checkNiceStatus, userID, ArticleID)
	if err := row.Scan(&niceID); err != nil {
		if err == sql.ErrNoRows {
			return status, nil
		}
		log.Println(err)
		return status, err
	}
	if niceID != "" {
		status = true
	}
	return status, nil
}