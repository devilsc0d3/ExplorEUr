//add post ou commentaire

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

//button for report

const sendDataReportPost = async (txt) => {

    await fetch('http://localhost:8080/info', {
        method: 'POST',
        body: new URLSearchParams({
            reportPostButton: txt,
        })
    });
}

const sendDataReportComment = async (txt) => {

    await fetch('http://localhost:8080/info', {
        method: 'POST',
        body: new URLSearchParams({
            reportCommentButton: txt,
        })
    });
}

//text inside when report

const sendTextPostReport = async (txt) => {
    await fetch('http://localhost:8080/info', {
        method: 'POST',
        body: new URLSearchParams({
            textPostReport: txt,
        })
    });
};

const sendTextCommentReport = async (txt) => {
    await fetch('http://localhost:8080/info', {
        method: 'POST',
        body: new URLSearchParams({
            textCommentReport: txt,
        })
    });
};

function newPost() {
    let form = document.createElement('form');
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
    form.appendChild(sub);

    form.addEventListener("submit", (event) => {
        event.preventDefault();
        let newDiv = document.createElement("div");
        newDiv.innerHTML = txt.value;
        newDiv.classList.add("posts");
        let formComment = document.createElement("form");
        let txtComment = document.createElement('textarea');
        txtComment.setAttribute("name", "comment")
        txtComment.className = 'form_comment';
        formComment.appendChild(txtComment);
        let subComment = document.createElement("button");
        subComment.type = "submit";
        subComment.name = "comment";
        subComment.innerHTML = "Comment";
        formComment.appendChild(subComment);
        newDiv.appendChild(formComment);
        form.insertAdjacentElement('afterend', newDiv);
        sendDataPost(txt.value).then(res => res.json());
        form.reset();
        formComment.addEventListener("submit", (event) => {
            event.preventDefault();
            let newDivComment = document.createElement("div");
            newDivComment.innerHTML = txtComment.value;
            newDivComment.classList.add('comment');
            let button1 = document.querySelectorAll("[name = 'reportCommentButton']")[i];
            newDivComment.appendChild(button1);
            button1.addEventListener("click", () => {
                let text = newDiv.querySelector("p");
                sendTextCommentReport(text.innerHTML).then(r => r != null);
                sendDataReportComment(" ").then(r => r != null);
            })
            const postId = formComment.parentNode.dataset.id;
            formComment.insertAdjacentElement('beforebegin', newDivComment);
            sendDataComment(txtComment.value, postId).then(res => res.json());
            formComment.reset();
        });
        let button3 = document.querySelectorAll("[name = 'reportPostButton']")[i];
        button3.addEventListener("click", () => {
            let text = newDiv[+1].querySelector("p");
            sendTextPostReport(text.innerHTML).then(r => r != null);
            sendDataReportPost(" ").then(r => r != null);
        })
    });
}

function oldPost() {
    let posts = document.getElementsByClassName("posts");
    for (let i = 0; i < posts.length; i++) {
        let form = document.createElement("form");
        let txtComment = document.createElement('textarea');
        txtComment.setAttribute("name", "comment")
        txtComment.className = 'form_comment';
        form.appendChild(txtComment);
        let subComment = document.createElement("button");
        subComment.type = "submit";
        subComment.name = "comment";
        subComment.innerHTML = "Comment";
        form.appendChild(subComment);
        posts[i].appendChild(form);

        form.addEventListener("submit", (event) => {
            event.preventDefault();
            let newDivComment2 = document.createElement("div");
            newDivComment2.innerHTML = txtComment.value;
            newDivComment2.classList.add('comment');
            let button2 = document.querySelectorAll("[name = 'reportCommentButton']")[i];
            newDivComment2.appendChild(button2);
            button2.addEventListener("click", () => {
                let text = posts[i + 1].querySelector("p");
                        sendTextCommentReport(text.innerHTML).then(r => r != null);
                sendDataReportComment(" ").then(r => r != null);
            })
            const postId = form.parentNode.dataset.id;
            form.insertAdjacentElement('beforebegin', newDivComment2);
            sendDataComment(txtComment.value, postId).then(res => res.json());
            form.reset();
        });
        let button = document.querySelectorAll("[name = 'reportPostButton']")[i];
        button.addEventListener("click", () => {
            let text = posts[i + 1].querySelector("p");
            sendTextPostReport(text.innerHTML).then(r => r != null);
            sendDataReportPost(" ").then(r => r != null);
        })
    }
}

export function chat() {
    oldPost();
    newPost();
}