import axios from "axios";

export async function shortenLink(links){
    let bodyFormData = new FormData()
    bodyFormData.append('url', links)


    return await axios({
        method: 'post',
        url: process.env.REACT_APP_URL,
        data: bodyFormData
    }).then(function (response) {
            return response.data
        }
    )
}