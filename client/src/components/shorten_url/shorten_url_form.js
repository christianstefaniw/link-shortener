import {Component} from 'react'
import {Form} from "reactstrap";

import ShortenUrlInput from "./shorten_url_input";
import ShortenUrlButton from "./shorten_url_button";

import {shortenLink} from "../../services/shorten_link"

class ShortenUrlForm extends Component {

    constructor(props) {
        super(props);
        this.state = {
            links: ''
        }
        this.inputChangeHandler = this.inputChangeHandler.bind(this)
        this.onSubmit = this.onSubmit.bind(this)
        this.props.onLinkChange.bind(this)
    }

    inputChangeHandler(e) {
        this.setState({links: e.target.value})
    }

    onSubmit() {
        shortenLink(this.state.links).then((result) => {
            this.props.onLinkChange(result)
        })
    }

    render() {
        return (
            <Form>
                <ShortenUrlInput inputChangeHandler={this.inputChangeHandler}/>
                <ShortenUrlButton onSubmit={this.onSubmit}/>
            </Form>
        );
    }
}

export default ShortenUrlForm