package models

import (
	_ "errors"
	"fmt"
	//"reflect"
	"github.com/astaxie/beego/orm"

	_ "strings"
)

type Post struct {
	Id       int    //
	User     *User  `orm:"rel(fk)"`
	Title    string //
	Subtitle string //
	Text     string `orm:"type(text)"` //
	Tags     []*Tag `orm:"rel(m2m)"`
}

func GetPostById(id int) (post Post, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable("post")
	err = qs.Filter("Id", id).RelatedSel().One(&post)
	if err == orm.ErrNoRows {
		err = ErrNoRows
	}
	return
}

func (this Post) Update() (num int64, err error) {
	o := orm.NewOrm()
	num, err = o.Update(&this)
	if err == orm.ErrNoRows {
		err = ErrNoRows
	}
	return
}

func (this *Post) Delete() (num int64, err error) {
	o := orm.NewOrm()
	num, err = o.Delete(this)
	if err == orm.ErrNoRows {
		err = ErrNoRows
	}
	return
}

func (this *Post) New() (num int64, err error) {
	o := orm.NewOrm()
	num, err = o.Insert(this)

	if err == orm.ErrNoRows {
		err = ErrNoRows
	}
	return
}

type stack []string

func (s *stack) Push(v string) {
	*s = append(*s, v)
}

func (s *stack) Pop() string {
	res := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return res
}

func isPostValidByTags(post_tags []string, query_tags []string) (result bool) {
	result = false
	var st stack
	fmt.Println("!!!!!!!!!!!!!!!!!")
	fmt.Println(post_tags)
	fmt.Println(query_tags)
	for _, s := range query_tags {
		fmt.Println("!")

		fmt.Println(post_tags)
		fmt.Println(st)
		if s != "*" && s != "+" {
			st.Push(s)
		} else {
			var count = 0
			to_push := "("
			for j := 0; j < 2; j++ {
				s_pop := st.Pop()
				fmt.Println(">", s_pop)
				to_push += s_pop + s
				fmt.Println("<", to_push)
				for _, s1 := range post_tags {
					fmt.Println(":", s1)
					if s1 == s_pop {
						count++
						break
					}

				}
				fmt.Println("=", count)

			}
			to_push = to_push[:len(to_push)-1]
			to_push += ")"
			st.Push(to_push)
			if (s == "*" && count >= 2) || (s == "+" && count >= 1) {
				post_tags = append(post_tags, to_push)
			}
		}
		fmt.Println(st)
	}

	result = false
	s_pop := st.Pop()
	for _, s1 := range post_tags {
		fmt.Println(">>>>>>>", s1)
		if s1 == s_pop {
			result = true
			break
		}
	}

	return
}

func GetPostsByTags(tags_str []string) (num int64, posts []Post, tags []Tag, err error) {

	var tags3 []string
	for _, s := range tags_str {
		if s != "*" && s != "+" {
			tags3 = append(tags3, s)
		}
	}
	_, posts_2, _ := GetPostsByAnyTags(tags3)
	type tags_count struct {
		Tag   *Tag
		Count uint
	}

	var tags_c []tags_count

	for _, p := range posts_2 {
		var tags2 []string
		for _, t := range p.Tags {
			tags2 = append(tags2, t.Name)
		}

		if isPostValidByTags(tags2, tags_str) == true {
			fmt.Println(":::novo post válido:::", p.Title)
			posts = append(posts, *p)

			// insere e ordena as tags deste post validado
			for i, t := range p.Tags {
				// skip some dafult tags
				if i < 2 {
					continue
				}
				fmt.Println("::nova tag válida::", t.Name)
				k := 0
				for j, tc := range tags_c {
					fmt.Println("j:", j)
					k = j
					if t.Name == tc.Tag.Name {
						fmt.Println("primeiro if")
						tags_c[j].Count++
						break
					} else if j == len(tags_c)-1 {
						fmt.Println("segundo if")
						k++
						break
					}
				}
				if k+1 > len(tags_c) {
					tags_c = append(tags_c, tags_count{Count: 1, Tag: t})
				}
				// ordena de k pra trás
				for j := k; j > 0; j-- {
					temp := tags_c[j-1]
					if temp.Count < tags_c[j].Count {
						tags_c[j-1] = tags_c[j]
						tags_c[j] = temp
					}
				}
				fmt.Println(":tags:")
				for _, t2 := range tags_c {
					fmt.Println(t2.Tag.Name, t2.Count)
				}
				fmt.Println("::::::")
			}

		}
	}
	for _, t := range tags_c {
		tags = append(tags, *t.Tag)
	}
	return
}

func GetPostsByAnyTags(tags []string) (num int64, posts []*Post, err error) {
	o := orm.NewOrm()

	params := make([]interface{}, 0)
	for _, t := range tags {
		params = append(params, t)
	}
	var posts_q []Post
	o.QueryTable("post").Filter("Tags__Tag__Name__in", params...).RelatedSel().Distinct().OrderBy("-id").All(&posts_q)

	for i, _ := range posts_q {
		posts = append(posts, &posts_q[i])
		o.LoadRelated(&posts_q[i], "Tags")
	}

	if err == orm.ErrNoRows {
		err = ErrNoRows
	}
	return
}
