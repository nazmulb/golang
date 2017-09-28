<html> 
    <head>
        <title>Provider Score</title>
        <link rel="stylesheet" href="/static/css/styles.css" />
    </head>
    <body>
        <table border="1">
            <th>ID</th>
            <th>Provider Score Rank</th>
            {{range .}}
                <tr>
                    <td>{{.ProviderUserID}}</td>
                    <td>{{.ProviderScoreRank}}</td>
                </tr>
            {{end}}
        </table>
    </body>
</html>