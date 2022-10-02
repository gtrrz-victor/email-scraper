const aws = require('aws-sdk');
const { simpleParser } =  require('mailparser')
const { JSDOM } = require('jsdom')
const s3 = new aws.S3({ apiVersion: '2006-03-01' });

exports.handler = async (event) => {
    const bucket = event.Records[0].s3.bucket.name;
    const key = decodeURIComponent(event.Records[0].s3.object.key.replace(/\+/g, ' '));
    const params = {
        Bucket: bucket,
        Key: key,
    };
    try {
        const { Body } = await s3.getObject(params).promise();
        const mail = await simpleParser(Body)
        const dom = new JSDOM(mail.html)
        const contactEmail = getContactEmail(dom.window.document)
        const phoneNumber = getPhoneNumber(dom.window.document)
        const numPeople = getNumberPeople(dom.window.document)
        return { contactEmail, phoneNumber, numPeople }
    } catch (err) {
        console.error(err);
        const message = `Error getting object ${key} from bucket ${bucket}. Make sure they exist and your bucket is in the same region as this function.`;
        throw new Error(message);
    }
};


function getContactEmail(document) {
  return document.querySelector('body > div > div > div:nth-child(4) > div > table:nth-child(6) > tbody > tr:nth-child(2) > td > table > tbody > tr > td > table > tbody > tr:nth-child(8) > td > table > tbody > tr:nth-child(7) > td > div > div > a').textContent
}

function getPhoneNumber(document) {
  return document.querySelector('body > div > div > div:nth-child(4) > div > table:nth-child(6) > tbody > tr:nth-child(2) > td > table > tbody > tr > td > table > tbody > tr:nth-child(8) > td > table > tbody > tr:nth-child(7) > td > div > strong:nth-child(10) > a').textContent
}

function getNumberPeople(document) {
  return document.querySelector('body > div > div > div:nth-child(4) > div > table:nth-child(6) > tbody > tr:nth-child(2) > td > table > tbody > tr > td > table > tbody > tr:nth-child(8) > td > table > tbody > tr:nth-child(10) > td > table > tbody > tr > td > table:nth-child(1) > tbody > tr:nth-child(1) > td:nth-child(1) > div').textContent
}