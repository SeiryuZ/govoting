
$('.upvote').on('click', function() {
	var id = $(this).data('id');
  var vote_item_id = $(this).data('vote-item-id');
	$.ajax({
	  type: "POST",
	  url: "/upvote",
	  data: { id: id, vote_item_id: vote_item_id },
    success: function() {
      var $upvote = $('.upvote[data-id="'+id+'"]');
      $upvote.find('i').addClass('active');
      $upvote_count = $upvote.parent().find('.upvote-count');
      $upvote_count.html(parseInt($upvote_count.html()) + 1);

    },
    error: function(xhr, message) {
      if (xhr.status == 403) {
        var url = xhr.responseText;
        window.location = url;
      }
      if (xhr.status == 400) {
        console.log("Already voted");
      }
    }
	});
});