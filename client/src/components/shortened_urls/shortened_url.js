import {Component} from 'react'
import {Col, Row} from "reactstrap";

import './shortened_url_style.css'

class ShortenedUrl extends Component{

    render() {
        return (
            <Row className="justify-content-between url-row" style={{backgroundColor: this.props.backgroundColor}}>
                <Col>
                    <p className="full-url">{this.props.link['FullURL']}</p>
                </Col>
                <Col>
                    <p className="short-url">{`${process.env.REACT_APP_URL}/${this.props.link['ShortURL']}`}</p>
                </Col>

            </Row>
        );
    }
}

export default ShortenedUrl