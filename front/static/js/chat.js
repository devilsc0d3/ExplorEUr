const sendDataPost = async (txt) => {

    await fetch('http://localhost:8080/info', {
        method: 'POST',
        body: new URLSearchParams({
            postContent: txt,
        })
    });
};

const sendDataComment = async (txt, postId) => {

    await fetch('http://localhost:8080/info', {
        method: 'POST',
        body: new URLSearchParams({
            comment: txt,
            postID: postId,
        })
    });
};

const sendDataLike = async (like, dislike, postId) => {

    await fetch('http://localhost:8080/info', {
        method: 'POST',
        body: new URLSearchParams({
            like: like,
            dislike: dislike,
            postId: postId
        })
    });
};

function newPost() {
    let like = false;
    let dislike = false;
    let countClickLike = 0;
    let countClickDislike = 0;

    //add input for new post
    let form = document.querySelector('[name="newPost"]')
    console.log(form)
    form.classList.add('posts');
    const header = document.querySelector("header")
    header.insertAdjacentElement('afterend', form);

    let txt = document.createElement('textarea');
    txt.setAttribute("name", "postContent")
    form.appendChild(txt);

    let sub = document.createElement("button");
    sub.type = "submit";
    sub.name = "post";
    sub.innerHTML = "Share";
    sub.className = "button"
    form.appendChild(sub);

    //event when submit your post
    form.addEventListener("submit", (event) => {
        event.preventDefault();

        //add post
        let newDiv = document.createElement("div");
        newDiv.innerHTML = txt.value;
        newDiv.classList.add("posts");

        //add comment input
        let formComment = document.createElement("form");
        let txtComment = document.createElement('textarea');
        txtComment.setAttribute("name", "comment")
        txtComment.className = 'form_comment';
        formComment.appendChild(txtComment);
        let subComment = document.createElement("button");
        subComment.type = "submit";
        subComment.name = "comment";
        subComment.innerHTML = "Comment";
        subComment.className = "button"
        formComment.appendChild(subComment);
        newDiv.appendChild(formComment);

        //add button like/dislike
        let buttonLike2 = document.createElement("button");
        buttonLike2.classList.add("like");
        newDiv.appendChild(buttonLike2);
        let imgLike = document.createElement("img");
        imgLike.src = "http://localhost:8080/static/image/likeFALSE.png";
        imgLike.classList.add("img_like");
        buttonLike2.appendChild(imgLike);
        let buttonDislike2 = document.createElement("button");
        buttonDislike2.classList.add("dislike");
        newDiv.appendChild(buttonDislike2);
        let imgDislike = document.createElement("img");
        imgDislike.src = "http://localhost:8080/static/image/dislikeFALSE.png";
        imgDislike.classList.add("img_like");
        buttonDislike2.appendChild(imgDislike);

        form.insertAdjacentElement('afterend', newDiv);
        sendDataPost(txt.value).then(res => res.json()).catch(res => Promise.fail({error: res}));
        form.reset();

        //event when submit your comment
        formComment.addEventListener("submit", (event) => {
            event.preventDefault();
            let newDivComment = document.createElement("div");
            newDivComment.innerHTML = txtComment.value;
            newDivComment.classList.add('comment');
            const postId = formComment.parentNode.dataset.id;
            formComment.insertAdjacentElement('beforebegin', newDivComment);
            sendDataComment(txtComment.value, postId).then(res => res.json());
            formComment.reset();
        });

        //event when like
        buttonLike2.addEventListener("click", (event2) => {
            countClickLike++;
            event2.preventDefault();
            if (!dislike) {
                if (countClickLike % 2 === 0) {
                    imgLike.src = "http://localhost:8080/static/image/likeFALSE.png"
                    like = false;
                } else {
                    imgLike.src = "http://localhost:8080/static/image/likeTRUE.png"
                    like = true;
                    const postId = formComment.parentNode.dataset.id;
                    sendDataLike(like, dislike, postId).then(r => r != null);
                }
            }
        });

        //event when dislike
        buttonDislike2.addEventListener("click", (event2) => {
            countClickDislike++;
            event2.preventDefault();
            if (!like) {
                if (countClickDislike % 2 === 0) {
                    imgDislike.src = "http://localhost:8080/static/image/dislikeFALSE.png"
                    dislike = false;
                } else {
                    imgDislike.src = "http://localhost:8080/static/image/dislikeTRUE.png"
                    dislike = true;
                    const postId = formComment.parentNode.dataset.id;
                    sendDataLike(like, dislike, postId).then(r => r != null);
                }
            }
        });
    });
}

function oldPost() {
    let like = false;
    let dislike = false;
    let countClickLike = 0;
    let countClickDislike = 0;
    let posts = document.getElementsByClassName("posts");

    for (let i = 0; i < posts.length; i++) {

        //add input for comment
        let form = posts[i].querySelector("[name='newComment']");
        let txtComment = document.createElement('textarea');
        txtComment.setAttribute("name", "comment")
        txtComment.className = 'form_comment';
        form.appendChild(txtComment);

        let subComment = document.createElement("button");
        subComment.type = "submit";
        subComment.name = "comment";
        subComment.innerHTML = "Comment";
        subComment.className = "button"
        form.appendChild(subComment);
        posts[i].appendChild(form);

        //add button like/dislike
        let buttonLike = document.createElement("button");
        buttonLike.classList.add("like");
        posts[i].appendChild(buttonLike);
        let imgLike = document.createElement("img");
        imgLike.src = "http://localhost:8080/static/image/likeFALSE.png";
        imgLike.classList.add("img_like");
        buttonLike.appendChild(imgLike);
        let countLike = posts[i].dataset.like;
        let nbrLike = document.createElement("h6");
        nbrLike.innerHTML = countLike;
        posts[i].appendChild(nbrLike)

        let buttonDislike = document.createElement("button");
        buttonDislike.classList.add("dislike");
        posts[i].appendChild(buttonDislike);
        let imgDislike = document.createElement("img");
        imgDislike.src = "http://localhost:8080/static/image/dislikeFALSE.png";
        imgDislike.classList.add("img_like");
        buttonDislike.appendChild(imgDislike);


        //event when submit your comment
        form.addEventListener("submit", (event) => {
            event.preventDefault();
            let newDivComment = document.createElement("div");
            newDivComment.innerHTML = txtComment.value;
            newDivComment.classList.add('comment');
            const postId = form.parentNode.dataset.id;
            form.insertAdjacentElement('beforebegin', newDivComment);

            sendDataComment(txtComment.value, postId).then(res => res.json());

            form.reset();
        });

        //event when like
        buttonLike.addEventListener("click", (event) => {
            countClickLike++;
            event.preventDefault();
            if (!dislike) {
                if (countClickLike % 2 === 0) {
                    imgLike.src = "http://localhost:8080/static/image/likeFALSE.png"
                    like = false;
                } else {
                    imgLike.src = "http://localhost:8080/static/image/likeTRUE.png"
                    like = true;
                    const postId = form.parentNode.dataset.id;
                    sendDataLike(like, dislike, postId).then(r => r!=null)
                }
            }
        });

        //event when dislike
        buttonDislike.addEventListener("click", (event) => {
            countClickDislike++;
            event.preventDefault();
            if (!like) {
                if (countClickDislike % 2 === 0) {
                    imgDislike.src = "http://localhost:8080/static/image/dislikeFALSE.png"
                    dislike = false;
                } else {
                    imgDislike.src = "http://localhost:8080/static/image/dislikeTRUE.png"
                    dislike = true;
                    const postId = form.parentNode.dataset.id;
                    sendDataLike(like, dislike, postId).then(r => r!=null)
                }
            }
        });
    }
}

export function chat() {
    oldPost();
    newPost();
}