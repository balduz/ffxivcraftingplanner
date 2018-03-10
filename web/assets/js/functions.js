addToCraftingListFunction = function(e) {
  const elem = $(e.target).parent().parent();

  if (elem) {
    const copy = elem.clone();
    copy.children().last().remove();
    copy.append('<div class="col text-left"><button type="button" class="btn btn-danger">Remove</button></div>');
    copy.children().last().click(removeFunction);
    $('#crafting-list').append(copy);

    const id = elem.data("id");
    $.get('/recipe/' + id, function(data) {
      $('#gathering-mats-list').append(data);

      $('.ingredients-modal').click(function() {
        $('#exampleModalCenter').modal();
      });
    });
  }
}

removeFunction = function(e) {
  const elem = $(e.target).parent().parent();
  elem.remove();
}

// Register click functions.
$('.btn-primary').on('click', addToCraftingListFunction);
$('#crafting-list > .btn-danger').on('click', removeFunction);

$('#item-search-form').on('submit', function(e) {
  // Prevent refresh of the whole page.
  e.preventDefault();

  var $input = $('#item-search-form :input').first();
  const val = $input.val();

  $('#search-results-list').parent().load('/search', {"search": val}, function() {
    $('#search-results-list .btn-primary').click(addToCraftingListFunction)
    // TODO(cbalduz): check whether I need to do something here or not, like an error message.
  });
});
