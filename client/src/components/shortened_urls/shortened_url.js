import {Component} from 'react'
import {Col, Row} from "reactstrap";

import './shortened_url_style.css'

class ShortenedUrl extends Component{

    render() {
        return (
            <Row className="justify-content-between url-row" style={{backgroundColor: this.props.backgroundColor}}>
                <Col>
                    <p className="shortened-item">{this.props.link['FullURL']}</p>
                </Col>
                <Col>
                    <p className="shortened-item">{`http://localhost:8080/${this.props.link['ShortURL']}`}</p>
                </Col>

            </Row>
        );
    }
}

export default ShortenedUrl