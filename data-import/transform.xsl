<?xml version="1.0" encoding="UTF-8"?>
<xsl:stylesheet version="3.0"
                xmlns:xsl="http://www.w3.org/1999/XSL/Transform"
                xmlns:xs="http://www.w3.org/2001/XMLSchema"
                xmlns:map="http://www.w3.org/2005/xpath-functions/map"
                xmlns:array="http://www.w3.org/2005/xpath-functions/array"
                exclude-result-prefixes="xs map array">

    <xsl:output method="json" indent="yes"/>

    <xsl:template match="/library">
        <xsl:sequence select="map{'library': array{book}}"/>
    </xsl:template>

    <xsl:template match="book">
        <xsl:sequence select="map{
         'id': string(@id),
         'title': title,
         'author': author
      }"/>
    </xsl:template>

</xsl:stylesheet>
