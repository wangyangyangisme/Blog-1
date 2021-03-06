{{template "admin/header.tpl" .}}

<div id="body">
<form class="entry" action="/admin/entries/{{.PostId}}" method="POST">
	<input class="entry-title" type="text" name="title" value="{{.Entry.Title}}" placeholder="Title">
	<input class="entry-subtitle" type="text" name="subtitle" value="{{.Entry.Subtitle}}" placeholder="Subtitle">
	<textarea class="entry-content-edit" type="text" name="content" placeholder="Content">{{.Entry.Content}}</textarea>
	<div class="entry-content"></div>

	<h1>{{.Entry.Collection}}</h1>
	<select name="collection" value="{{.Entry.Collection}}">
		{{if eq .Entry.Collection ""}}
			<option value="" selected>None</option>
		{{else}}
			<option value="">None</option>
		{{end}}

		{{range $index, $c := .Collections}}
			{{if eq $c.Title .Entry.Collection}}
				<option value="{{$c.Title}}" selected>{{$c.Title}}</option>
			{{else}}
				<option value="{{$c.Title}}">{{$c.Title}}</option>
			{{end}}
		{{end}}
	</select>
	<input class="post-button" type="submit" value="POST">
	<div class="msg">{{.Message}}</div>
</form>
</div>

{{template "admin/footer.tpl" .}}