import axios from "axios";
import {getBaseURL} from "./constants";

const fs = require("fs").promises;

const relativePath = "/essay"
const url = getBaseURL() + relativePath

describe("Inserts and manipulates essays", () => {

    it("Inserts multiple essays",  async () => {

        for(const i in ["a", "b", "c", "d"]) {
            const q = {
                "title": i,
                "text": "abcdef",
                "topicId": "1"
            }
            const resp = await axios.post(url, q)
            expect(resp.status).toBe(201)
        }
        const resp = await axios.post(url, {
            "title":  "e",
            "text": "ghi",
            "topicId": "2"
        })
        expect(resp.status).toBe(201)
    })

    it("Retrieve random essays by topic", async () => {
        const resp = await  axios.get(url + "?topicId=1&amount=10")
        const data = resp.data
        expect(resp.status).toBe(200)
        expect(data.data.length).toBe(5)
        for(let essay of data.data) {
            expect(essay.topicId).toEqual("1")
        }
    })

    it("Retrieve random essays", async () => {
        const resp = await  axios.get(url + "?mode=random&amount=3")
        const data = resp.data
        expect(resp.status).toBe(200)
        expect(data.data.length).toBe(3)
    })

    it("Retrieve essay using id", async () => {
        const initialResp = await  axios.get(url + "?mode=random")
        const initialData = initialResp.data
        const id = initialData.data[0].id
        const title = initialData.data[0].title

        const resp = await axios.get(url + "/" + id)
        const data = resp.data
        expect(resp.status).toBe(200)
        expect(data.data.title).toEqual(title)
    })

    it("Comment on an essay", async () => {
        const retrievalResp = await  axios.get(url + "?mode=random")
        const retrievalData = retrievalResp.data
        const id = retrievalData.data[0].id
        const title = retrievalData.data[0].title

        const commentResp = await axios.post(url + "/" + id + "/comment", {
            "text": "first"
        })
        expect(commentResp.status).toEqual(201)

        const resp = await axios.get(url + "/" + id)
        const data = resp.data.data
        expect(resp.status).toBe(200)
        expect(data.title).toEqual(title)
        expect(data.comments.length).toEqual(1)
        expect(data.comments[0]).toEqual("first")

        const secondCommentResp = await axios.post(url + "/" + id + "/comment", {
            "text": "second"
        })
        expect(secondCommentResp.status).toEqual(201)

        const secondResp = await axios.get(url + "/" + id)
        const secondData = secondResp.data.data
        expect(secondResp.status).toBe(200)
        expect(secondData.title).toEqual(title)
        expect(secondData.comments.length).toEqual(2)
        expect(secondData.comments[0]).toEqual("first")
        expect(secondData.comments[1]).toEqual("second")
    })

})
