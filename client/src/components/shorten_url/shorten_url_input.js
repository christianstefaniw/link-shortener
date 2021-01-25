import {Component} from 'react'
import {Input} from "reactstrap";

import "./shorten_url_style.css"

class ShortenUrlInput extends Component{

    constructor(props) {
        super(props);
        this.props.inputChangeHandler.bind(this)
    }

    render() {
        return (
            <Input style={{height: '150px'}} onChange={this.props.inputChangeHandler} type="textarea" name="links" placeholder="Shorten your links, separate each link with a comma (,)" />
        )
    }
}

export default ShortenUrlInput