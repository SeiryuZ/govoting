
$('.upvote').on('click', function() {
	var id = $(this).data('id');
	$.ajax({
	  type: "POST",
	  url: "/upvote",
	  data: { id: id }
	});
});