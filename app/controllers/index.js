import Ember from 'ember';

export default Ember.Controller.extend({
    sortKey: 'code',
    sortedItems: Ember.computed.sort('model', 'itemsSorting'),
    itemsSorting: Ember.computed(function () {
        return [this.get('sortKey')];
    })
});
