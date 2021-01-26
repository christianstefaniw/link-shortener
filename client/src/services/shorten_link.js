import axios from "axios";

export async function shortenLink(links){
    let bodyFormData = new FormData()
    bodyFormData.append('url', links)


    return await axios({
        method: 'post',
        url: 'http://192.168.1.131:8080',
        data: bodyFormData
    }).then(function (response) {
            return response.data
        }
    )
}