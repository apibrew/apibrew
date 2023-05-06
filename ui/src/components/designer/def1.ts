import {dia} from "jointjs";
import {Resource} from "../../model";




export class TheClass2 extends dia.Element {
    resource: Resource

    abc = 'taleh'

    defaults() {
        return {
            ...super.defaults,
            type: 'uml.Class',
            attrs: {
                rect: {'width': 200},

                '.uml-class-name-rect': {'stroke': 'black', 'stroke-width': 2, 'fill': '#3498db'},
                '.uml-class-attrs-rect': {'stroke': 'black', 'stroke-width': 2, 'fill': '#2980b9'},
                '.uml-class-methods-rect': {'stroke': 'black', 'stroke-width': 2, 'fill': '#2980b9'},

                '.uml-class-name-text': {
                    'ref': '.uml-class-name-rect',
                    'ref-y': .5,
                    'ref-x': .5,
                    'text-anchor': 'middle',
                    'y-alignment': 'middle',
                    'font-weight': 'bold',
                    'fill': 'black',
                    'font-size': 12,
                    'font-family': 'Times New Roman'
                },
                '.uml-class-attrs-text': {
                    'ref': '.uml-class-attrs-rect', 'ref-y': 5, 'ref-x': 5,
                    'fill': 'black', 'font-size': 12, 'font-family': 'Times New Roman'
                },
                '.uml-class-methods-text': {
                    'ref': '.uml-class-methods-rect', 'ref-y': 5, 'ref-x': 5,
                    'fill': 'black', 'font-size': 12, 'font-family': 'Times New Roman'
                },
                '.uml-class-props-text': {
                    'ref': '.uml-class-methods-rect', 'ref-y': 5, 'ref-x': 5,
                    'fill': 'black', 'font-size': 12, 'font-family': 'Times New Roman'
                }
            },
            markup: [
                `<g class="rotatable"></g>`
            ].join('')
        }
    }

    redrawResource() {
        this.prop('markup', `<g class="rotatable">
                    <g class="scalable">
                        <rect class="uml-class-name-rect"/>
                        <rect class="uml-class-props-rect"/>
                        <rect class="uml-class-methods-rect"/>
                    </g>
                    <text class="uml-class-name-text">${this.resource.name}</text>
                    <text class="uml-class-props-text">${this.resource.name}</text>
                    ${this.resource.properties?.map((property, index) =>
            `<text class="prop-${index}">${property.name}</text>`)}
                </g>`)

        this.attr('.prop-0', {
            'ref': '.uml-class-methods-rect', 'ref-y': 50, 'ref-x': 50,
            'fill': 'black', 'font-size': 12, 'font-family': 'Times New Roman'
        })
        this.attr('.prop-1', {
            'ref': '.uml-class-methods-rect', 'ref-y': 50, 'ref-x': 150,
            'fill': 'black', 'font-size': 12, 'font-family': 'Times New Roman'
        })
    }

    constructor(resource: Resource) {
        super();

        this.resource = resource

        console.log(this)
    }

    initialize() {
        this.on('change:name change:attributes change:methods', function () {
            // @ts-ignore
            this.updateRectangles();
            // @ts-ignore
            this.trigger('uml-update');
            console.log('triggered')
        }, this);

        super.initialize()
    }

    updateText() {
        this.abc = 'updated abc'

        this.trigger('update')
    }
}


